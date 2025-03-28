package main

import (
	"context"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	lookup "github.com/gagliardetto/solana-go/programs/address-lookup-table"
	"github.com/gagliardetto/solana-go/rpc"
	"golang.org/x/time/rate"
)

func main() {
	cluster := rpc.DevNet_RPC

	rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
		cluster,
		rate.Every(time.Second), // time frame
		5,                       // limit of requests per time frame
	))

	version := uint64(0)
	tx, err := rpcClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("48n7jpRoDgUU12638z779XNeyzFSR5GwFMcjRUTkerhj7WtAiMDAgrYuH8DzJ2rEhXAUdvLD58Z1S5PB1N933YJe"),
		&rpc.GetTransactionOpts{
			MaxSupportedTransactionVersion: &version,
			Encoding:                       solana.EncodingBase64,
		},
	)
	if err != nil {
		panic(err)
	}
	parsed, err := tx.Transaction.GetTransaction()
	if err != nil {
		panic(err)
	}
	fmt.Println("parsed tx: %v", parsed)
	//processTransactionWithAddressLookups(parsed, rpcClient)
}

func processTransactionWithAddressLookups(txx *solana.Transaction, rpcClient *rpc.Client) {
	if !txx.Message.IsVersioned() {
		fmt.Println("tx is not versioned; only versioned transactions can contain lookups")
		return
	}
	tblKeys := txx.Message.GetAddressTableLookups().GetTableIDs()
	if len(tblKeys) == 0 {
		fmt.Println("no lookup tables in versioned transaction")
		return
	}
	numLookups := txx.Message.GetAddressTableLookups().NumLookups()
	if numLookups == 0 {
		fmt.Println("no lookups in versioned transaction")
		return
	}
	fmt.Println("num lookups:", numLookups)
	fmt.Println("num tbl keys:", len(tblKeys))
	resolutions := make(map[solana.PublicKey]solana.PublicKeySlice)
	for _, key := range tblKeys {
		fmt.Println("Getting table", key)

		info, err := rpcClient.GetAccountInfo(
			context.Background(),
			key,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("got table " + key.String())

		tableContent, err := lookup.DecodeAddressLookupTableState(info.GetBinary())
		if err != nil {
			panic(err)
		}

		fmt.Println("table content:", spew.Sdump(tableContent))
		fmt.Println("isActive", tableContent.IsActive())

		resolutions[key] = tableContent.Addresses
	}

	err := txx.Message.SetAddressTables(resolutions)
	if err != nil {
		panic(err)
	}

	err = txx.Message.ResolveLookups()
	if err != nil {
		panic(err)
	}
	fmt.Println(txx.String())
}
