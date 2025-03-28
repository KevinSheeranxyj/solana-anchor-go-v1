package main

import (
	"context"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {

	pubKey := solana.MustPublicKeyFromBase58("6cGS4LoW9PA9C3ktMRRX6JN7ebQi9gKpEaqA7gYZ1mBr")
	client := rpc.New(rpc.DevNet_RPC)
	resp, err := client.GetAccountInfo(
		context.TODO(),
		pubKey,
	)
	if err != nil {
		panic(err)
	}

	borshDec := bin.NewBorshDecoder(resp.GetBinary())
	var meta
	err = borshDec.Decode(&meta)
	if err != nil {
		panic(err)
	}
}
