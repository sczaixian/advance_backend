package main

import (
	"advance_backend/test"
	"fmt"
)

func demo_solana() {
	//solana_test.TestAddress()
	//solana_test.TestCreateAccount()
	//solana_test.ExampleFromBase64()
	//solana_test.TestCreateAccount()
	//solana_test.TestWalletTransaction()
	//solana_test.TestGetAccountInfo()
	//solana_test.TestRpcGetBalance()
	//solana_test.TestRpcGetBlock()  // TODO:
	//solana_test.TestRpcGetCommitment()
	//solana_test.TestRpcGetBlockHeight()
	//solana_test.TestRpcGetBlockProduction()
	//solana_test.TestRpcGetBlockTime()
	//solana_test.TestRpcGetBlocks()
	//solana_test.TestRpcGetBlockWithLimit()
	//solana_test.TestRpcGetClusterNodes()

}

func demo_test() {
	client := test.TestClient()
	_ = client
	fmt.Println("-------client-----------\n\n")
	clientLocal := test.TestClientLocal()
	_ = clientLocal
	fmt.Println("-------clientLocal-----------\n\n")
	client_ws := test.TestClientWebSocket()
	_ = client_ws
	fmt.Println("-------clientWebSocket--------\n\n")
	address := test.TestAddress(client)
	fmt.Println("-------TestAddress-----------\n\n")
	test.TestBlock(client)
	fmt.Println("-------TestBlock-----------\n\n")
	test.TestBalance(client, &address)
	fmt.Println("-------TestBalance-----------\n\n")
	//test.SelectTransfer(client)
	//fmt.Println("-------SelectTransfer-----------\n\n")
	//test.SelectReceipt(client)
	//fmt.Println("-------SelectReceipt-----------\n\n")
	//test.GenPrivateKey()
	//fmt.Println("-------GenPrivateKey-----------\n\n")
	////test.TestTransfer(client)
	//fmt.Println("-------TestTransfer-----------\n\n")
	//test.SelectTokenBalance(client)
	//fmt.Println("-------SelectTokenBalance-----------\n\n")
	//test.TestTokenTransfer(client)
	//fmt.Println("-------TestTokenTransfer-----------\n\n")
	//test.TestSubscribe(client_ws)
	//fmt.Println("-------TestSubscribe-----------\n\n")
}

func self_test() {
	//client := self_test.SepoliaClient()
	//_ = client
	//client_ws := self_test.WebSocketClinet()
	//_ = client_ws
	//
	//self_test.TestAddress(client)

}

func main() {
	//demo_test()
	demo_solana()
}
