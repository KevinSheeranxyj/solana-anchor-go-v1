package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
	ctx := context.Background()
	client, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)
	}
	program := solana.MustPublicKeyFromBase58("6cGS4LoW9PA9C3ktMRRX6JN7ebQi9gKpEaqA7gYZ1mBr") // serum

	{
		// Subscribe to log events that mention the provided pubkey:
		sub, err := client.LogsSubscribeMentions(
			program,
			rpc.CommitmentConfirmed,
		)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()

		for {
			got, err := sub.Recv(ctx)
			if err != nil {
				panic(err)
			}
			spew.Dump(got)
		}
	}
	if false {
		// Subscribe to all log events:
		sub, err := client.LogsSubscribe(
			ws.LogsSubscribeFilterAll,
			rpc.CommitmentConfirmed,
		)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()

		for {
			got, err := sub.Recv(ctx)
			if err != nil {
				panic(err)
			}
			spew.Dump(got)
		}
	}
}
