package rpc_client_v2

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
)

func Stake() {
	// Fetch and decode the SupernodeState
	//state, err := fetchSupernodeState(rpcClient, supernode)
	//if err != nil {
	//	log.Fatalf("Error fetching SupernodeState: %v", err)
	//}
	//
	//// Print the deserialized data
	//fmt.Printf("Admin: %s\n", state.Admin.String())
	//fmt.Printf("Token: %s\n", state.Token.String())
	//fmt.Printf("Decimals: %d\n", state.Decimals)
	//fmt.Printf("Reward Locked Time: %d\n", state.RewardLockedTime)
	//fmt.Printf("Staking Coefficient: %d\n", state.StakingCoefficient)
	//fmt.Printf("K Values: %v\n", state.KValues)
	//
	//// Example: Get k_value for spec_id 0
	//kValue, err := state.GetDeviceKValue(deviceId)
	//if err != nil {
	//	log.Fatalf("Error getting k_value: %v", err)
	//}
	//amount := float64(kValue) / 1_000_000_000
	//fmt.Printf("K Value for spec_id %d: amount: %.9f, actual k value %d\n", deviceId, amount, kValue)
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
	// Deserialize using borsh
	var state SupernodeState
	err = borsh.Deserialize(&state, data)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize account data: %v", err)
	}

	return &state, nil
}

type SupernodeState struct {
	Admin              solana.PublicKey `json:"admin"`
	Token              solana.PublicKey `json:"token"`
	Decimals           uint64           `json:"decimals"`
	RewardLockedTime   uint64           `json:"reward_locked_time"`
	StakingCoefficient uint64           `json:"stake_coefficient"`
	KValues            [100]uint64      `json:"k_values"`
}

func (s *SupernodeState) GetDeviceKValue(specID uint64) (uint64, error) {
	if specID >= 100 {
		return 0, fmt.Errorf("spec_id %d out of bounds (max 99)", specID)
	}
	return s.KValues[specID], nil
}
