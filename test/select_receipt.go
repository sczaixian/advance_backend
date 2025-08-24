package test

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// SelectReceipt 函数演示了如何通过以太坊客户端查询区块和交易回执信息
func SelectReceipt() {
	// 创建与以太坊节点的连接
	// URL 是以太坊节点的RPC端点地址（需提前定义）
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err) // 连接失败时终止程序
	}

	// 创建一个表示特定区块号的大整数（BLOCK_NUMBER 需提前定义）
	blockNumber := big.NewInt(BLOCK_NUMBER)
	// 将十六进制字符串转换为哈希类型（用于指定区块哈希）
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")

	// 通过区块哈希获取指定区块的所有交易回执
	receiptByHash, err := client.BlockReceipts(
		context.Background(),
		rpc.BlockNumberOrHashWithHash(blockHash, false), // false 表示不要求完整交易对象
	)
	if err != nil {
		log.Fatal(err)
	}

	// 通过区块号获取同一区块的所有交易回执
	receiptsByNum, err := client.BlockReceipts(
		context.Background(),
		rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 验证两种方式获取的第一个回执是否相同（应返回 true）
	fmt.Println(receiptByHash[0] == receiptsByNum[0])

	// 遍历通过哈希获取的回执列表（此处仅处理第一个回执后跳出循环）
	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)                // 交易状态：1 表示成功
		fmt.Println(receipt.Logs)                  // 交易触发的日志事件（空数组表示无事件）
		fmt.Println(receipt.TxHash.Hex())          // 交易哈希的十六进制表示
		fmt.Println(receipt.TransactionIndex)      // 交易在区块中的索引位置
		fmt.Println(receipt.ContractAddress.Hex()) // 合约地址（为零地址表示非合约创建交易）
		break                                      // 仅处理第一个回执后退出循环
	}

	// 直接通过交易哈希获取特定交易的收据
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c7c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	// 打印该交易的详细信息
	fmt.Println(receipt.Status)                // 交易状态
	fmt.Println(receipt.Logs)                  // 日志事件
	fmt.Println(receipt.TxHash.Hex())          // 交易哈希
	fmt.Println(receipt.TransactionIndex)      // 交易索引
	fmt.Println(receipt.ContractAddress.Hex()) // 合约地址
}

//func SelectReceipt() {
//	client, err := ethclient.Dial(URL)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	blockNumber := big.NewInt(BLOCK_NUMBER)
//	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
//	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true
//
//	for _, receipt := range receiptByHash {
//		fmt.Println(receipt.Status)                // 1
//		fmt.Println(receipt.Logs)                  // []
//		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
//		fmt.Println(receipt.TransactionIndex)      // 0
//		fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
//		break
//	}
//
//	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
//	receipt, err := client.TransactionReceipt(context.Background(), txHash)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(receipt.Status)                // 1
//	fmt.Println(receipt.Logs)                  // []
//	fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
//	fmt.Println(receipt.TransactionIndex)      // 0
//	fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
//}
