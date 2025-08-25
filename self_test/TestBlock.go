package self_test

import (
	"context"
	"fmt"
)


func TestBlock(client * ethclient.Client){
	blockNumber := big.NewInt(block_height)
	header, err := client.HeaderByNumber(context.BackGround(), blockNumber)
	if err != nill {
		fmt.Fatal(".xx")
	}

	fmt.Println(header.Number.Uint64(), header.Number.String())
	fmt.Println(header.Time)
	fmt.Println(header.Difficulty.Uint64())
	fmt.Println(header.Hash().Hex())

	block, err := client.BlockByNumber(context.BackGround(), blockNumber)
	block.Number().Uint64()
	block.Time()
	block.Difficulty().Uint64()
	block.Hash().Hex()
	block.Transactions()

	client.TransactionCount(context.Background(), block.Hash())
	
}