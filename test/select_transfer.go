package test

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SelectTransfer(client *ethclient.Client) {

	// 2. 获取当前网络的链ID
	// 链ID用于EIP-155签名，防止交易在不同链上重放
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 3. 指定要查询的区块号
	blockNumber := big.NewInt(BLOCK_HEIGHT)

	// 4. 通过区块号获取完整的区块信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 5. 遍历区块中的所有交易
	for _, tx := range block.Transactions() {
		fmt.Println("交易哈希:", tx.Hash().Hex())          // 打印交易哈希（交易的唯一标识符）
		fmt.Println("交易金额(wei):", tx.Value().String()) // 打印交易金额（以wei为单位）
		fmt.Println("Gas限制:", tx.Gas())                  // 打印Gas限制（执行交易所需的最大计算工作量）
		fmt.Println("Gas价格:", tx.GasPrice().Uint64())    // 打印Gas价格（每个Gas单位的价格，以wei为单位）
		fmt.Println("Nonce:", tx.Nonce())                  // 打印发送方的nonce值（交易计数器）
		fmt.Println("交易数据:", tx.Data())                // 打印交易数据（对于普通转账为空，合约调用包含调用数据）
		// 打印接收方地址（对于合约创建交易，此值为nil）
		if tx.To() != nil {
			fmt.Println("接收方地址:", tx.To().Hex()) // 打印接收方地址
		} else {
			fmt.Println("合约创建交易") // 地址为nil合约为创建交易
		}

		// 6. 使用EIP-155签名器从交易中提取发送方地址
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("发送方地址:", sender.Hex())
		} else {
			log.Fatal(err)
		}

		// 7. 获取交易回执（包含交易执行结果）
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		// 打印交易状态（1表示成功，0表示失败）
		fmt.Println("交易状态:", receipt.Status)

		// 打印交易日志（智能合约事件）
		fmt.Println("交易日志:", receipt.Logs)

		// 只处理第一笔交易后跳出循环
		break
	}

	// 8. 通过区块哈希获取区块中的交易数量
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	fmt.Println("交易数量: ", count)
	if err != nil {
		log.Fatal(err)
	}

	// 9. 遍历区块中的所有交易（通过索引）
	for idx := uint(0); idx < count; idx++ {
		// 通过区块哈希和交易索引获取交易
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		// 打印交易哈希
		fmt.Println("交易哈希:", tx.Hash().Hex())

		// 只处理第一笔交易后跳出循环
		break
	}

	// 10. 通过交易哈希直接获取交易
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	// 打印交易是否仍在待处理状态（内存池中）
	fmt.Println("是否待处理:", isPending)

	// 打印交易哈希
	fmt.Println("交易哈希:", tx.Hash().Hex())

	// 再次打印是否待处理状态
	fmt.Println("是否待处理:", isPending) // false 表示已确认}
}
