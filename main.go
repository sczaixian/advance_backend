package main

import (
	"advance_backend/alchemy_api"
	"advance_backend/eth_test"
	"advance_backend/opensea_api"
	"advance_backend/solana_and_eth"
	"advance_backend/solana_test"
	"advance_backend/tasks"
	"advance_backend/test"
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
	ethClient := solana_and_eth.EthClient()
	solanaClient := solana_and_eth.SolanaClientTestNet()
	fmt.Println("-------init client-----------\n\n")
	ethBlock := solana_and_eth.EthSelectBlock(ethClient)
	fmt.Println("-------  EthSelectBlock  -----------\n\n")
	solanaBlock := solana_and_eth.SolanaSelectBlock(solanaClient)
	fmt.Println("-------  SolanaSelectBlock  -----------\n\n")
	solana_and_eth.EthTransaction(ethClient, ethBlock)
	fmt.Println("-------  EthTransaction  -----------\n\n")
	solana_and_eth.SolanaTransaction(solanaClient, solanaBlock)
	fmt.Println("-------  SolanaTransaction  -----------\n\n")
	// TODO: eth 转账， 代币转账
	solana_and_eth.EthBalance(ethClient)
	fmt.Println("-------  EthBalance  -----------\n\n")
	solana_and_eth.SolanaBalance()
	fmt.Println("-------  SolanaBalance  -----------\n\n")
	solana_and_eth.EthSubscribes()
	fmt.Println("-------  EthSubscribes  -----------\n\n")
	solana_and_eth.SolanaSubscribesAccount()
	fmt.Println("-------  SolanaSubscribesAccount  -----------\n\n")
	solana_and_eth.SolanaSubscribesLog()
	fmt.Println("-------  SolanaSubscribesLog  -----------\n\n")
	solana_and_eth.SolanaSubscribesSol()
	fmt.Println("-------  SolanaSubscribesSol  -----------\n\n")

}

func gen_wallet() {
	solana_and_eth.EthWalletCreate()
	fmt.Println("-------  EthWalletCreate  -----------\n\n")
	solana_and_eth.SolanaWalletCreate()
	fmt.Println("-------  SolanaWalletCreate  -----------\n\n")
	solana_and_eth.EthWalletCreateDeepseek()
	fmt.Println("-------  EthWalletCreateDeepseek  -----------\n\n")
	solana_and_eth.SolanaWalletCreateDeepseek()
	fmt.Println("-------  SolanaWalletCreateDeepseek  -----------\n\n")
}

func eth_test_rlp() {
	eth_test.Test1()
}

func tasksExec() {
	//blockNumber := big.NewInt(9249114)
	//tasks.Task1(blockNumber)
	//fmt.Println("-------  Task1  -----------\n\n")
	//tasks.Task1Tx()
	//fmt.Println("-------  Task1Tx  -----------\n\n")
	//tasks.CounterTest()
	//fmt.Println("-------  CounterTest  -----------\n\n")
	//tasks.Task3Solana()
	tasks.Task3Transfer()
}

func openseaAPI() {
	// docs :  https://docs.opensea.io/reference/api-keys
	apiKey := "6098d14358814514a0cae4d5369bb77f"
	account_address := "0x3E0bDb54f94D735dDCf8D2074c852a8C22914aA7"
	c_address := "0x6571a7350f0f879f89f7f434ea32ec4c1f947df7"
	perfix := "https://api.opensea.io/api/v2/"
	sepolia_chain_id := "11155111"
	urls := make(map[string]string)
	//get account: https://api.opensea.io/api/v2/accounts/{address_or_username}
	urls["get_account"] = "%s/accounts/%s"
	//getNFT:  https://api.opensea.io/api/v2/chain/{chain}/contract/{address}/nfts/{identifier}
	urls["get_nfg"] = "%s/chain/%s/contract/%s/nfts/%s"
	//by account: https://api.opensea.io/api/v2/chain/{chain}/account/{address}/nfts
	urls["account"] = "%s/chain/%s/account/%s/nfts"
	//by collection: https://api.opensea.io/api/v2/collection/{collection_slug}/nfts
	//by contract: https://api.opensea.io/api/v2/chain/{chain}/contract/{address}/nfts
	//get payment token: https://api.opensea.io/api/v2/chain/{chain}/payment_token/{address}
	//opensea_api.AccountAPI()
	//opensea_api.CollectionAPI()
	//opensea_api.GetCollectionsAPI()
	opensea_api.API_Call(fmt.Sprintf(urls["get_account"], perfix, account_address), apiKey)
	opensea_api.API_Call(fmt.Sprintf(urls["account"], perfix, sepolia_chain_id, c_address), apiKey)
}

func alchemyAPI() {
	//alchemy_api.API_Call()
	alchemy_api.AlchemyAPI()
}

func main() {
	//demo_test()
	//demo_solana()
	//self_test()
	//gen_wallet()
	//check_demo.TestXXXBBBB()

	//eth_test_rlp()
	//tasksExec()
	//openseaAPI()
	alchemyAPI()
}
