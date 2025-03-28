package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
)

// ProviderStakeInfo matches the Rust account structure
type ProviderStakeInfo struct {
	ExtraControllers [2]solana.PublicKey // 2 Pubkeys (32 bytes each)
	Devices          []DeviceState       // Vec<DeviceState>
}

// DeviceState matches the Rust struct
type DeviceState struct {
	State              uint16
	SpecID             uint16
	StakingCoefficient uint64
	KValue             uint64
}

// CheckOldControllerExists checks if old_controller exists in provider_stake_info.extra_controllers
func CheckOldControllerExists(
	ctx context.Context,
	rpcClient *rpc.Client,
	programID solana.PublicKey,
	provider solana.PublicKey,
	oldController solana.PublicKey,
) (bool, error) {
	// Derive the provider_stake_info PDA
	seeds := [][]byte{
		[]byte("provider_stake_info"),
		provider.Bytes(),
	}
	pda, _, err := solana.FindProgramAddress(seeds, programID)
	if err != nil {
		return false, fmt.Errorf("failed to derive PDA: %v", err)
	}

	// Fetch account data
	accountInfo, err := rpcClient.GetAccountInfo(ctx, pda)
	if err != nil {
		return false, fmt.Errorf("failed to fetch account %s: %v", pda, err)
	}
	if accountInfo == nil || len(accountInfo.Value.Data.GetBinary()) == 0 {
		return false, fmt.Errorf("account %s not found or empty", pda)
	}

	// Deserialize the account data (skip 8-byte Anchor discriminator)
	var providerStakeInfo ProviderStakeInfo
	err = borsh.Deserialize(&providerStakeInfo, accountInfo.Value.Data.GetBinary()[8:])
	if err != nil {
		return false, fmt.Errorf("failed to deserialize ProviderStakeInfo: %v", err)
	}

	// Check if old_controller exists in extra_controllers
	for _, controller := range providerStakeInfo.ExtraControllers {
		spew.Dump(controller)
		if controller.Equals(oldController) {
			return true, nil
		}
	}
	return false, nil
}

func main() {
	// Initialize RPC client (Devnet)
	rpcClient := rpc.New(rpc.DevNet_RPC)
	ctx := context.Background()

	// Example keys (replace with real values)
	programID := solana.MustPublicKeyFromBase58("6cGS4LoW9PA9C3ktMRRX6JN7ebQi9gKpEaqA7gYZ1mBr")     // Your program ID
	provider := solana.MustPublicKeyFromBase58("Fnv6Bsc1QQd8umNpZgbvtuwjjBEjtSPzFHuHFdczUaJx")      // Provider's pubkey
	oldController := solana.MustPublicKeyFromBase58("Fnv6Bsc1QQd8umNpZgbvtuwjjBEjtSPzFHuHFdczUaJx") // Old controller to check

	// Check if old_controller exists
	exists, err := CheckOldControllerExists(ctx, rpcClient, programID, provider, oldController)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if exists {
		fmt.Printf("Old controller %s exists in extra_controllers for provider %s\n", oldController, provider)
	} else {
		fmt.Printf("Old controller %s does not exist in extra_controllers for provider %s\n", oldController, provider)
	}
}
