package solana_test

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func TestGetAccountInfo() {
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)

	{
		pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
		// basic usage
		resp, err := client.GetAccountInfo(
			context.TODO(),
			pubKey,
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(resp)

		var mint token.Mint
		// Account{}.Data.GetBinary() returns the *decoded* binary data
		// regardless the original encoding (it can handle them all).

		//TODO:
		//err = bin.NewDecoder(resp.GetBinary()).Decode(&mint)
		//
		//if err != nil {
		//	panic(err)
		//}

		spew.Dump(mint)
		// NOTE: The supply is mint.Supply, with the mint.Decimals:
		// mint.Supply = 9998022451607088
		// mint.Decimals = 6
		// ... which means that the supply is 9998022451.607088
	}
	{
		// Or you can use `GetAccountDataInto` which does all of the above in one call:
		pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
		var mint token.Mint
		// Get the account, and decode its data into the provided mint object:
		err := client.GetAccountDataInto(
			context.TODO(),
			pubKey,
			&mint,
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(mint)
	}
	{
		// // Or you can use `GetAccountDataBorshInto` which does all of the above in one call but for borsh-encoded data:
		// var metadata token_metadata.Metadata
		// // Get the account, and decode its data into the provided metadata object:
		// err := client.GetAccountDataBorshInto(
		//   context.TODO(),
		//   pubKey,
		//   &metadata,
		// )
		// if err != nil {
		//   panic(err)
		// }
		// spew.Dump(metadata)
	}
	{
		pubKey := solana.MustPublicKeyFromBase58("4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R") // raydium token
		// advanced usage
		resp, err := client.GetAccountInfoWithOpts(
			context.TODO(),
			pubKey,
			// You can specify more options here:
			&rpc.GetAccountInfoOpts{
				Encoding:   solana.EncodingBase64Zstd,
				Commitment: rpc.CommitmentFinalized,
				// You can get just a part of the account data by specify a DataSlice:
				// DataSlice: &rpc.DataSlice{
				//  Offset: pointer.ToUint64(0),
				//  Length: pointer.ToUint64(1024),
				// },
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(resp)

		var mint token.Mint
		//TODO:
		//err = bin.NewDecoder(resp.GetBinary()).Decode(&mint)
		//if err != nil {
		//	panic(err)
		//}
		spew.Dump(mint)
	}
}
