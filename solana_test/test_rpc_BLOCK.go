package solana_test

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func TestRpcGetBlock() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	example, err := client.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	{
		out, err := client.GetBlock(context.TODO(), uint64(example.Context.Slot))
		if err != nil {
			panic(err)
		}
		// spew.Dump(out) // NOTE: This generates a lot of output.
		spew.Dump(len(out.Transactions))
	}

	{
		includeRewards := false
		out, err := client.GetBlockWithOpts(
			context.TODO(),
			uint64(example.Context.Slot),
			// You can specify more options here:
			&rpc.GetBlockOpts{
				Encoding:   solana.EncodingBase64,
				Commitment: rpc.CommitmentFinalized,
				// Get only signatures:
				TransactionDetails: rpc.TransactionDetailsSignatures,
				// Exclude rewards:
				Rewards: &includeRewards,
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
}

func TestRpcGetBlockHeight() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	out, err := client.GetBlockHeight(
		context.TODO(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}

func TestRpcGetBlockProduction() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	{
		out, err := client.GetBlockProduction(context.TODO())
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
	{
		out, err := client.GetBlockProductionWithOpts(
			context.TODO(),
			&rpc.GetBlockProductionOpts{
				Commitment: rpc.CommitmentFinalized,
				// Range: &rpc.SlotRangeRequest{
				//  FirstSlot: XXXXXX,
				//  Identity:  solana.MustPublicKeyFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				// },
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
}

func TestRpcGetBlockTime() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	example, err := client.GetLatestBlockhash(
		context.TODO(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}

	out, err := client.GetBlockTime(
		context.TODO(),
		uint64(example.Context.Slot),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
	spew.Dump(out.Time().Format(time.RFC1123))
}

func TestRpcGetBlockWithLimit() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	example, err := client.GetLatestBlockhash(
		context.TODO(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}

	limit := uint64(4)
	out, err := client.GetBlocksWithLimit(
		context.TODO(),
		uint64(example.Context.Slot-10),
		limit,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}

func TestRpcGetBlocks() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	example, err := client.GetLatestBlockhash(
		context.TODO(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}

	endSlot := uint64(example.Context.Slot)
	out, err := client.GetBlocks(
		context.TODO(),
		uint64(example.Context.Slot-3),
		&endSlot,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
