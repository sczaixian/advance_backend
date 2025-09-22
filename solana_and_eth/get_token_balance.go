package solana_and_eth

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func EthTokenBalance(client *ethclient.Client) {
	fmt.Println("todo:....")
}

func SolanaTokenBalance(client *rpc.Client) {
	pubKey := solana.MustPublicKeyFromBase58("BJgJSEj1fqjC8uabfMw2Pho4sAt3YJbMcWuVkKNddaSn")
	out, err := client.GetTokenAccountBalance(context.TODO(), pubKey, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)

	//pubKey = solana.MustPublicKeyFromBase58("AfkALUPjQp8R1rUwE6KhT38NuTYWCncwwHwcJu7UtAfV")
	//out, err := client.GetTokenAccountsByDelegate(
	//	context.TODO(),
	//	pubKey,
	//	&rpc.GetTokenAccountsConfig{
	//		Mint: solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112"),
	//	},
	//	nil)
	//
	//if err != nil {
	//	panic(err)
	//}
	//spew.Dump(out)
}
