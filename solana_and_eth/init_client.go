package solana_and_eth

import (
	"advance_backend/test"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gagliardetto/solana-go/rpc"
	"golang.org/x/time/rate"
)

func EthClient() *ethclient.Client {
	client, err := ethclient.Dial(test.URL) //  根据url，可以是https 或 websocket
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func EthClientWS() *ethclient.Client {
	client, err := ethclient.Dial(test.WS_URL) //  根据url，可以是https 或 websocket
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func SolanaClientMainNet() *rpc.Client {
	cluster := rpc.MainNetBeta
	rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
		cluster.RPC, rate.Every(time.Second), 5))
	return rpcClient
}

func SolanaClientDevNet() *rpc.Client {
	endpoint := rpc.DevNet_RPC
	client := rpc.New(endpoint)
	return client
}

func SolanaClientTestNet() *rpc.Client {
	//cluster := rpc.MainNetBeta
	//cluster := rpc.DevNet
	//cluster := rpc.TestNet
	//rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
	//	cluster.RPC, rate.Every(time.Second), 5))
	//return rpcClient

	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)
	return client
}
