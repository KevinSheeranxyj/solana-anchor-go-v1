package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	logging "github.com/ipfs/go-log/v2"
	"github.com/near/borsh-go"
)

var log = logging.Logger("unstake-amount-extraction")

// DeviceUnstakeEvent matches the Rust event structure
type DeviceUnstakeEvent struct {
	Provider               solana.PublicKey // 32 bytes
	DeviceID               uint64           // 8 bytes
	Amount                 uint64           // 8 bytes
	ProviderVestingInfoKey solana.PublicKey // 32 bytes
}

func GetUnstakeAmount(ctx context.Context, rpcClient *rpc.Client, txSig solana.Signature) (uint64, error) {
	// Fetch transaction
	tx, err := rpcClient.GetTransaction(ctx, txSig, &rpc.GetTransactionOpts{
		Encoding:   solana.EncodingBase64,
		Commitment: rpc.CommitmentConfirmed,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to fetch tx %s: %v", txSig, err)
	}
	if tx == nil || tx.Transaction == nil || tx.Meta == nil {
		return 0, fmt.Errorf("transaction %s not found or incomplete", txSig)
	}

	// Confirm it's an unstake_device instruction
	solTx, err := tx.Transaction.GetTransaction()
	if err != nil {
		return 0, fmt.Errorf("failed to decode tx %s: %v", txSig, err)
	}
	unstakeDiscriminator := [8]byte{0x56, 0x50, 0x3C, 0x26, 0x48, 0xB9, 0xFA, 0x3D}
	found := false

	for _, inst := range solTx.Message.Instructions {
		if len(inst.Data) < 8 || inst.ProgramIDIndex >= uint16(len(solTx.Message.AccountKeys)) {
			continue
		}
		dataPrefix := [8]byte{}
		copy(dataPrefix[:], inst.Data[:8])
		fmt.Printf("Data: %v\n", inst.Data[:8])
		if dataPrefix == unstakeDiscriminator {
			found = true
			break
		}
	}
	if !found {
		return 0, fmt.Errorf("no unstake_device instruction in tx %s", txSig)
	}

	// Compute event discriminator
	eventHash := sha256.Sum256([]byte("DeviceUnstakeEvent"))
	eventDiscriminator := eventHash[:8] // First 8 bytes

	log.Debugf("log message: %v", tx.Meta.LogMessages)

	// Search logs for DeviceUnstakeEvent
	for _, logMsg := range tx.Meta.LogMessages {
		// Look for "Program data:" followed by base64-encoded event data
		if len(logMsg) < len("Program data: ") || logMsg[:len("Program data: ")] != "Program data: " {
			continue
		}
		base64Data := logMsg[len("Program data: "):]
		data, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			continue
		}

		// Check if it’s the event by matching discriminator
		if len(data) < 8 || !bytes.Equal(data[:8], eventDiscriminator) {
			continue
		}

		// Deserialize the event
		var event DeviceUnstakeEvent
		err = borsh.Deserialize(&event, data[8:]) // Skip discriminator
		if err != nil {
			log.Errorf("Failed to deserialize event in tx %s: %v", txSig, err)
			continue
		}

		log.Infof("Unstake Amount: %d lamports, DeviceID: %d", event.Amount, event.DeviceID)
		return event.Amount, nil
	}

	return 0, fmt.Errorf("DeviceUnstakeEvent not found in tx %s", txSig)
}

func main() {
	rpcClient := rpc.New("https://api.devnet.solana.com")
	ctx := context.Background()

	// Replace with a real unstake_device transaction signature
	txSig, _ := solana.SignatureFromBase58("48n7jpRoDgUU12638z779XNeyzFSR5GwFMcjRUTkerhj7WtAiMDAgrYuH8DzJ2rEhXAUdvLD58Z1S5PB1N933YJe")
	amount, err := GetUnstakeAmount(ctx, rpcClient, txSig)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Unstake Amount: %d lamports\n", amount)
}
