package self_test

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


func SelectTransfer(client * ethclient.Client){


	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(xxx)
	block, err := client.BlockByNumber(context.Background(), blockNumber)



	for _, tx := range block.Transactions(){
		tx.Hash().Hex()
		tx.Value().String()
		tx.Gas()
		tx.GasPrice().Uint64()
		tx.Nonce()
		tx.Data()
		tx.To().Hex()
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil{
			sender.Hex()
		}else{
			log.Fatal(err)
		}

		receipt, err := client.TransectionReceipt(context.Background(), tx.Hash())
		receipt.Status
		receipt.Logs

	}

	blockHash := common.HexToHash("")
	count, err := client.TransactionCount(context.Background(), blockHash)
	
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)

		tx.Hash().Hex()
	}

	txHash := common.HexToHash("addr")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	
	
}