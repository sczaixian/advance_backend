package main

import (
	"advance_backend/solana_test"
	"advance_backend/test"
	"advance_backend/xxx"
	"fmt"
)

func demo_solana() {
	solana_test.TestAddress()
	solana_test.TestCreateAccount()
	solana_test.ExampleFromBase64()
	solana_test.TestCreateAccount()
	solana_test.TestWalletTransaction()
	solana_test.TestGetAccountInfo()
	solana_test.TestRpcGetBalance()
	//solana_test.TestRpcGetBlock()  // TODO:
	solana_test.TestRpcGetCommitment()
	solana_test.TestRpcGetBlockHeight()
	solana_test.TestRpcGetBlockProduction()
	solana_test.TestRpcGetBlockTime()
	solana_test.TestRpcGetBlocks()
	solana_test.TestRpcGetBlockWithLimit()
	solana_test.TestRpcGetClusterNodes()

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
	test.SelectTransfer(client)
	fmt.Println("-------SelectTransfer-----------\n\n")
	test.SelectReceipt(client)
	fmt.Println("-------SelectReceipt-----------\n\n")
	test.GenPrivateKey()
	fmt.Println("-------GenPrivateKey-----------\n\n")
	//test.TestTransfer(client)
	fmt.Println("-------TestTransfer-----------\n\n")
	test.SelectTokenBalance(client)
	fmt.Println("-------SelectTokenBalance-----------\n\n")
	test.TestTokenTransfer(client)
	fmt.Println("-------TestTokenTransfer-----------\n\n")
	test.TestSubscribe(client_ws)
	fmt.Println("-------TestSubscribe-----------\n\n")
}

func self_test() {
	ethClient := xxx.EthClient()
	solanaClient := xxx.SolanaClientTestNet()
	fmt.Println("-------init client-----------\n\n")
	ethBlock := xxx.EthSelectBlock(ethClient)
	fmt.Println("-------  EthSelectBlock  -----------\n\n")
	solanaBlock := xxx.SolanaSelectBlock(solanaClient)
	fmt.Println("-------  SolanaSelectBlock  -----------\n\n")
	xxx.EthTransaction(ethClient, ethBlock)
	fmt.Println("-------  EthTransaction  -----------\n\n")
	xxx.SolanaTransaction(solanaClient, solanaBlock)
	fmt.Println("-------  SolanaTransaction  -----------\n\n")
	// TODO: eth 转账， 代币转账
	xxx.EthBalance(ethClient)
	fmt.Println("-------  EthBalance  -----------\n\n")
	xxx.SolanaBalance()
	fmt.Println("-------  SolanaBalance  -----------\n\n")
	xxx.EthSubscribes()
	fmt.Println("-------  EthSubscribes  -----------\n\n")
	xxx.SolanaSubscribesAccount()
	fmt.Println("-------  SolanaSubscribesAccount  -----------\n\n")
	xxx.SolanaSubscribesLog()
	fmt.Println("-------  SolanaSubscribesLog  -----------\n\n")
	xxx.SolanaSubscribesSol()
	fmt.Println("-------  SolanaSubscribesSol  -----------\n\n")

}

func gen_wallet() {
	xxx.EthWalletCreate()
	fmt.Println("-------  EthWalletCreate  -----------\n\n")
	xxx.SolanaWalletCreate()
	fmt.Println("-------  SolanaWalletCreate  -----------\n\n")
	xxx.EthWalletCreateDeepseek()
	fmt.Println("-------  EthWalletCreateDeepseek  -----------\n\n")
	xxx.SolanaWalletCreateDeepseek()
	fmt.Println("-------  SolanaWalletCreateDeepseek  -----------\n\n")
}

func main() {
	//demo_test()
	//demo_solana()
	self_test()
	//gen_wallet()
	//check_demo.TestXXXBBBB()
}
