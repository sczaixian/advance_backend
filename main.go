package main

import (
	"advance_backend/test"
	"fmt"
)

func main() {
	client := test.TestClient()
	_ = client
	fmt.Println("-------client-----------\n\n")
	clientLocal := test.TestClientLocal()
	_ = clientLocal
	fmt.Println("-------clientLocal-----------\n\n")
	client_ws := test.TestClientWebSocket()
	_ = client_ws
	fmt.Println("-------clientWebSocket--------\n\n")
	//address := test.TestAddress(client)
	//fmt.Println("-------TestAddress-----------\n\n")
	//test.TestBlock(client)
	//fmt.Println("-------TestBlock-----------\n\n")
	//test.TestBalance(client, &address)
	//fmt.Println("-------TestBalance-----------\n\n")
	//test.SelectTransfer(client)
	//fmt.Println("-------SelectTransfer-----------\n\n")
	//test.SelectReceipt(client)
	//fmt.Println("-------SelectReceipt-----------\n\n")
	//test.GenPrivateKey()
	//fmt.Println("-------GenPrivateKey-----------\n\n")
	////test.TestTransfer(client)
	//fmt.Println("-------TestTransfer-----------\n\n")
	test.SelectTokenBalance(client)
	//fmt.Println("-------SelectTokenBalance-----------\n\n")
	//test.TestTokenTransfer(client)
	//fmt.Println("-------TestTokenTransfer-----------\n\n")
	//test.TestSubscribe(client_ws)
	//fmt.Println("-------TestSubscribe-----------\n\n")
}
