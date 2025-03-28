package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/near/borsh-go"
	"log"
	"strings"
)

// ProviderControllerChangedEvent matches the Rust event structure
type ProviderControllerChangedEvent struct {
	Provider      solana.PublicKey // 32 bytes
	Action        string           // variable-length string
	NewController solana.PublicKey // 32 bytes
	Operator      solana.PublicKey // 32 bytes
	OldController solana.PublicKey // 32 bytes
}

// MonitorReplaceControllerEvents monitors the ReplaceExtraController events via WebSocket
func MonitorReplaceControllerEvents(programID solana.PublicKey) {
	// Connect to Solana Devnet RPC and WebSocket
	rpcClient := rpc.New(rpc.DevNet_RPC)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer wsClient.Close()

	// Subscribe to program logs
	sub, err := wsClient.LogsSubscribe(
		ws.LogsSubscribeFilterAll,
		rpc.CommitmentConfirmed,
	)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}
	defer sub.Unsubscribe()

	// Compute event discriminator for ProviderControllerChangedEvent
	eventHash := sha256.Sum256([]byte("ProviderControllerChangedEvent"))
	eventDiscriminator := eventHash[:8]

	// Instruction discriminator for replace_extra_controller (sha256("global:replace_extra_controller")[0:8])
	instructionHash := sha256.Sum256([]byte("global:replace_extra_controller"))
	instructionDiscriminator := instructionHash[:8]

	fmt.Printf("Monitoring program %s for ReplaceExtraController events...\n", programID)

	// Process incoming logs
	for {
		result, err := sub.Recv(context.Background())
		if err != nil {
			log.Printf("Error receiving logs: %v", err)
			continue
		}

		// Check each log entry
		for _, logMsg := range result.Value.Logs {
			// Look for "Program data:" logs containing event data
			if strings.HasPrefix(logMsg, "Program data: ") {
				base58Data := strings.TrimPrefix(logMsg, "Program data: ")
				data, err := solana.PublicKeyFromBase58(base58Data)
				if err != nil {
					continue
				}

				// Check if it’s the event we’re looking for
				if len(data) < 8 || !bytes.Equal(data[:8], eventDiscriminator) {
					continue
				}

				// Deserialize the event (skip 8-byte discriminator)
				var event ProviderControllerChangedEvent
				err = borsh.Deserialize(&event, data[8:])
				if err != nil {
					log.Printf("Failed to deserialize event: %v", err)
					continue
				}

				// Verify it’s a "replace" action
				if event.Action == "replace" {
					fmt.Printf("ReplaceExtraController Event Detected:\n")
					fmt.Printf("  Provider: %s\n", event.Provider)
					fmt.Printf("  Action: %s\n", event.Action)
					fmt.Printf("  New Controller: %s\n", event.NewController)
					fmt.Printf("  Old Controller: %s\n", event.OldController)
					fmt.Printf("  Operator: %s\n", event.Operator)
					fmt.Printf("  Transaction Signature: %s\n", result.Value.Signature)
				}
			}
		}

		// Optionally, verify instruction discriminator in the transaction
		tx, err := rpcClient.GetTransaction(context.Background(), result.Value.Signature, &rpc.GetTransactionOpts{
			Encoding: solana.EncodingBase64,
		})
		if err != nil {
			continue
		}
		solTx, err := tx.Transaction.GetTransaction()
		if err != nil {
			continue
		}
		for _, inst := range solTx.Message.Instructions {
			if len(inst.Data) >= 8 && bytes.Equal(inst.Data[:8], instructionDiscriminator) {
				fmt.Println("Confirmed: replace_extra_controller instruction executed")
				break
			}
		}
	}
}

func main() {
	// Replace with your program ID
	programID := solana.MustPublicKeyFromBase58("6cGS4LoW9PA9C3ktMRRX6JN7ebQi9gKpEaqA7gYZ1mBr") // e.g., from `solana address`
	MonitorReplaceControllerEvents(programID)
}
