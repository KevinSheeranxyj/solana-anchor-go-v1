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

var log = logging.Logger("event-checker")

// DeviceStakedEvent matches the Rust struct
type DeviceStakedEvent struct {
	Provider solana.PublicKey // 32 bytes
	DeviceID uint64           // 8 bytes
	SpecID   uint64           // 8 bytes
	Amount   uint64           // 8 bytes
}

// DeviceUnstakeEvent matches the Rust struct
type DeviceUnstakeEvent struct {
	Provider               solana.PublicKey // 32 bytes
	DeviceID               uint64           // 8 bytes
	Amount                 uint64           // 8 bytes
	ProviderVestingInfoKey solana.PublicKey // 32 bytes
}

func CheckEventEmission(ctx context.Context, rpcClient *rpc.Client, txSig solana.Signature, eventName string) (bool, error) {
	// Fetch transaction
	tx, err := rpcClient.GetTransaction(ctx, txSig, &rpc.GetTransactionOpts{
		Encoding:   solana.EncodingBase64,
		Commitment: rpc.CommitmentConfirmed,
	})
	if err != nil {
		return false, fmt.Errorf("failed to fetch tx %s: %v", txSig, err)
	}
	if tx == nil || tx.Meta == nil {
		return false, fmt.Errorf("transaction %s not found or incomplete", txSig)
	}

	// Compute event discriminator
	eventHash := sha256.Sum256([]byte(eventName))
	eventDiscriminator := eventHash[:8]

	// Check logs for the event
	for _, logMsg := range tx.Meta.LogMessages {
		if len(logMsg) < len("Program data: ") || logMsg[:len("Program data: ")] != "Program data: " {
			continue
		}
		base58Data := logMsg[len("Program data: "):]
		fmt.Printf("base58Data: %s\n", base58Data)
		data, err := base64.StdEncoding.DecodeString(base58Data)
		if err != nil {
			continue
		}

		if len(data) < 8 || !bytes.Equal(data[:8], eventDiscriminator) {
			continue
		}

		// Event found, optionally deserialize to verify
		switch eventName {
		case "DeviceStakedEvent":
			var event DeviceStakedEvent
			if err := borsh.Deserialize(&event, data[8:]); err == nil {
				log.Infof("DeviceStakedEvent emitted - DeviceID: %d, Amount: %d", event.DeviceID, event.Amount)
				return true, nil
			}
		case "DeviceUnstakeEvent":
			var event DeviceUnstakeEvent
			if err := borsh.Deserialize(&event, data[8:]); err == nil {
				log.Infof("DeviceUnstakeEvent emitted - DeviceID: %d, Amount: %d", event.DeviceID, event.Amount)
				return true, nil
			}
		}
	}

	log.Infof("Event %s not found in tx %s", eventName, txSig)
	return false, nil
}

func main() {
	rpcClient := rpc.New("https://api.devnet.solana.com")
	ctx := context.Background()

	// Replace with your transaction signature
	txSig, _ := solana.SignatureFromBase58("48n7jpRoDgUU12638z779XNeyzFSR5GwFMcjRUTkerhj7WtAiMDAgrYuH8DzJ2rEhXAUdvLD58Z1S5PB1N933YJe")

	// Check for DeviceUnstakeEvent
	emitted, err := CheckEventEmission(ctx, rpcClient, txSig, "DeviceUnstakeEvent")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("DeviceUnstakeEvent emitted: %v\n", emitted)

	// Check for DeviceStakedEvent
	emitted, err = CheckEventEmission(ctx, rpcClient, txSig, "DeviceStakedEvent")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("DeviceStakedEvent emitted: %v\n", emitted)
}
