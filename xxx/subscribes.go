package xxx

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func EthSubscribes() {
	client := EthClientWS()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case headers := <-headers:
			spew.Dump(headers)
			fmt.Println("--------------- EthSubscribes headers ------------------")
			block, err := client.BlockByHash(context.Background(), headers.Hash())
			if err != nil {
				log.Fatal(err)
			}
			spew.Dump(block)
			fmt.Println("--------------- EthSubscribes block ------------------")
		}
	}
}

func SolanaSubscribesSol() {
	ctx := context.Background()
	client, err := ws.Connect(context.Background(), rpc.TestNet_WS)
	if err != nil {
		log.Fatal(err)
	}
	sub, err := client.SlotSubscribe()
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		got, err := sub.Recv(ctx)
		if err != nil {
			log.Fatal(err)
		}
		spew.Dump(got)
		fmt.Println("--------------- SolanaSubscribesSol --------------")
	}
}

func SolanaSubscribesAccount() {
	ctx := context.Background()
	client, err := ws.Connect(context.Background(), rpc.TestNet_WS)
	if err != nil {
		log.Fatal(err)
	}
	program := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
	{
		//sub, err := client.AccountSubscribe(program, "")
		//solana.EncodingJSONParsed  solana.EncodingJSON
		sub, err := client.AccountSubscribeWithOpts(program, "", solana.EncodingBase64)
		if err != nil {
			log.Fatal(err)
		}
		defer sub.Unsubscribe()
		for {
			got, err := sub.Recv(ctx)
			if err != nil {
				log.Fatal(err)
			}
			spew.Dump(got)
		}
	}
}

func SolanaSubscribesLog() {
	ctx := context.Background()
	client, err := ws.Connect(context.Background(), rpc.TestNet_WS)
	if err != nil {
		log.Fatal(err)
	}
	program := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
	{
		//  只订阅 特定账户
		sub, err := client.LogsSubscribeMentions(program, rpc.CommitmentRecent)
		// 订阅所有日志实践
		//sub, err := client.LogsSubscribe(ws.LogsSubscribeFilterAll,  rpc.CommitmentRecent)

		if err != nil {
			log.Fatal(err)
		}
		defer sub.Unsubscribe()
		for {
			got, err := sub.Recv(ctx)
			if err != nil {
				log.Fatal(err)
			}
			spew.Dump(got)
		}
	}
}
