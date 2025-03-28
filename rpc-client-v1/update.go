package main

import (
	"context"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
	"log"
)

// SupernodeState mirrors the Rust struct
type SupernodeState struct {
	Admin              solana.PublicKey `json:"admin"`
	Token              solana.PublicKey `json:"token"`
	Decimals           uint64           `json:"decimals"`
	RewardLockedTime   uint64           `json:"reward_locked_time"`
	StakingCoefficient uint64           `json:"stake_coefficient"`
	KValues            [100]uint64      `json:"k_values"`
}

func getSOLBalance(rpcClient *rpc.Client, walletAddress solana.PublicKey) (uint64, error) {
	// Fetch balance
	balance, err := rpcClient.GetBalance(
		context.Background(),
		walletAddress,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get SOL balance: %v", err)
	}

	return balance.Value, nil
}

func fetchMintTokenInfo(rpcClient *rpc.Client, pubKey solana.PublicKey) (*token.Mint, error) {
	resp, err := rpcClient.GetAccountInfoWithOpts(
		context.Background(),
		pubKey,
		&rpc.GetAccountInfoOpts{
			Encoding:   solana.EncodingBase64Zstd,
			Commitment: rpc.CommitmentFinalized,
		},
	)
	if err != nil {
		panic(err)
	}

	var mint token.Mint
	err = bin.NewBinDecoder(resp.GetBinary()).Decode(&mint)
	if err != nil {
		panic(err)
	}
	return &mint, nil
}

func GetATABalance(rpcClient *rpc.Client, walletAddress, tokenMint solana.PublicKey) (uint64, uint8, error) {
	// Step 1: Derive the ATA address
	ata, _, err := solana.FindAssociatedTokenAddress(walletAddress, tokenMint)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to derive ATA: %v", err)
	}

	// Step 2: Fetch ATA account info
	accountInfo, err := rpcClient.GetAccountInfoWithOpts(
		context.Background(),
		ata,
		&rpc.GetAccountInfoOpts{
			Commitment: rpc.CommitmentFinalized,
		},
	)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get ATA info: %v", err)
	}

	// Check if ATA exists
	if accountInfo.Value == nil {
		return 0, 0, fmt.Errorf("ATA %s does not exist or has no data", ata)
	}

	// Step 3: Decode the token account data
	var tokenAccount token.Account
	err = bin.NewBinDecoder(accountInfo.GetBinary()).Decode(&tokenAccount)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to deserialize token account: %v", err)
	}

	// Step 4: Fetch token mint decimals
	mintInfo, err := rpcClient.GetAccountInfoWithOpts(
		context.Background(),
		tokenMint,
		&rpc.GetAccountInfoOpts{
			Commitment: rpc.CommitmentFinalized,
		},
	)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get mint info: %v", err)
	}

	var mint token.Mint
	err = bin.NewBinDecoder(mintInfo.GetBinary()).Decode(&mint)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to deserialize mint: %v", err)
	}

	// Return the raw amount and decimals
	return tokenAccount.Amount, mint.Decimals, nil
}

func (s *SupernodeState) GetDeviceKValue(specID uint16) (uint64, error) {
	if specID >= 100 {
		return 0, fmt.Errorf("spec_id %d out of bounds (max 99)", specID)
	}
	return s.KValues[specID], nil
}

func fetchSupernodeState(rpcClient *rpc.Client, accountAddress solana.PublicKey) (*SupernodeState, error) {
	// Fetch account info
	accountInfo, err := rpcClient.GetAccountInfoWithOpts(
		context.Background(),
		accountAddress,
		&rpc.GetAccountInfoOpts{
			Commitment: rpc.CommitmentFinalized,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account info: %v", err)
	}

	// Check if account exists and has data
	if accountInfo.Value == nil {
		return nil, fmt.Errorf("account %s has no data", accountAddress)
	}

	// Get raw data (skip 8-byte discriminator for Anchor accounts)
	data := accountInfo.Value.Data.GetBinary()[8:] // Anchor accounts have an 8-byte discriminator prefix
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal account info: %v", err)
	//}
	// Deserialize using borsh
	var state SupernodeState
	err = borsh.Deserialize(&state, data)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize account data: %v", err)
	}

	return &state, nil
}

func main() {
	// Initialize RPC client (replace with your RPC endpoint)
	rpcClient := rpc.New("https://api.mainnet-beta.solana.com")

	// Replace with your SupernodeState account address
	accountAddress, err := solana.PublicKeyFromBase58("CvcBVZqNgSmsGMCUL76NKkpTZWgqrQomeAARmZs2C3Qg")
	if err != nil {
		log.Fatalf("Invalid account address: %v", err)
	}

	// Fetch and decode the SupernodeState
	state, err := fetchSupernodeState(rpcClient, accountAddress)
	if err != nil {
		log.Fatalf("Error fetching SupernodeState: %v", err)
	}

	// Print the deserialized data
	fmt.Printf("Admin: %s\n", state.Admin.String())
	fmt.Printf("Token: %s\n", state.Token.String())
	fmt.Printf("Decimals: %d\n", state.Decimals)
	fmt.Printf("Reward Locked Time: %d\n", state.RewardLockedTime)
	fmt.Printf("Staking Coefficient: %d\n", state.StakingCoefficient)
	fmt.Printf("K Values: %v\n", state.KValues)
	//
	//// Example: Get k_value for spec_id 0
	kValue, err := state.GetDeviceKValue(28)
	if err != nil {
		log.Fatalf("Error getting k_value: %v", err)
	}
	amount := float64(kValue) / 1_000_000_000
	fmt.Printf("K Value for spec_id 28: %.9f k value: (%d ) \n", amount, kValue)

	fmt.Printf("=============================: \n")

	mintAddr, err := solana.PublicKeyFromBase58("HRoehi5oj2L3qixJwRgEhpFSg2kzNvqJhH3B6zEXDN82")
	if err != nil {
		log.Fatalf("Error fetching SupernodeState: %v", err)
	}

	mint, err := fetchMintTokenInfo(rpcClient, mintAddr)
	if err != nil {
		log.Fatalf("Error fetching MintTokenInfo: %v", err)
	}
	fmt.Printf("Token: %v\n ", mintAddr)
	fmt.Printf("Supply: %d\n ", mint.Supply)
	fmt.Printf("Decimals: %d\n ", mint.Decimals)

	fmt.Printf("==================================\n")

	walletAddr, err := solana.PublicKeyFromBase58("9KwDdinWjgq7YZbGJf1JDH8p9MaT5XAda38wTqbbgGmo")
	if err != nil {
		log.Fatalf("Error fetching WalletAddress: %v", err)
	}

	balance, err := getSOLBalance(rpcClient, walletAddr)
	if err != nil {
		log.Fatalf("Error fetching Balance: %v", err)
	}
	solBalance := float64(balance) / 1_000_000_000
	fmt.Printf("Wallet address: %v\n", walletAddr)
	fmt.Printf("SOL Balance: %.9f SOL (%d lamports)\n", solBalance, balance)

	fmt.Printf("===========================\n")
	ataBalance, decimal, err := GetATABalance(rpcClient, walletAddr, mintAddr)
	if err != nil {
		log.Fatalf("Error fetching ATA Balance: %v", err)
	}
	afterbalance := float64(ataBalance) / 1_000_000_000
	fmt.Printf("ATA balance:  %.9f \t  raw balance: %d decimal: %d\n", afterbalance, ataBalance, decimal)

	result := 100001000000000/1_000_000_000 > 10054694634028/1_000_000_000
	fmt.Printf("Result: %v\n", result)
}
