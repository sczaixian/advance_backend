package self_test

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"advance_backend/test"
)

func SepoliaClient() {
	client, err := ethclient.Dial(test.URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("...")

	return client
}


func WebSocketClinet(){
	client, err := ethclient.Dial(test.WS_URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("...")
	return client
}

