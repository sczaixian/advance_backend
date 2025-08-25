package self_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)



func SelectReceipt(client *ethclient.Client){ 
	blockNumber := big.NewInt(111)
	blockHash := common.HexToHash("")

	receiptByHash, err := client.BlockReceipts(
		context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false),
	)

	receiptByNumber, err := client.BlockReceipts(
		context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.NlockNumber(blockNumber.Int64())),
	)

	for _, receipt := range receiptByHash {
		receipt.Status
		receipt.Logs
		receipt.TxHash.Hex()
		receipt.TransactionIndex
		receipt.ContractAddress.Hex()
	}

	txHash := common.HexToHash("")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)

	// ...
}