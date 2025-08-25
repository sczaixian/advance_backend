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
func SelectReceipt(client *ethclient.Client) {

	// 创建一个表示特定区块号的大整数（BLOCK_NUMBER 需提前定义）
	blockNumber := big.NewInt(BLOCK_HEIGHT) // 区块高度
	// 将十六进制字符串转换为哈希类型（用于指定区块哈希）
	blockHash := common.HexToHash(BLOCK_HASH) // 区块hash

	// 通过区块哈希获取指定区块的所有交易回执
	receiptByHash, err := client.BlockReceipts(
		context.Background(),
		rpc.BlockNumberOrHashWithHash(blockHash, false), // false 表示不要求完整交易对象
	)
	fmt.Println("receiptByHash:", receiptByHash)
	fmt.Println("err:", err)
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
	fmt.Println(receiptsByNum)

	// 验证两种方式获取的第一个回执是否相同（应返回 true）
	fmt.Println("验证是否相等：", receiptByHash[0] == receiptsByNum[0])

	for _, receipt := range receiptByHash {
		fmt.Println("交易状态：", receipt.Status)                // 交易状态：1 表示成功
		fmt.Println("日志事件:", receipt.Logs)                  // 交易触发的日志事件（空数组表示无事件）
		fmt.Println("交易哈希:", receipt.TxHash.Hex())          // 交易哈希的十六进制表示
		fmt.Println("在区块中的索引位置:", receipt.TransactionIndex) // 交易在区块中的索引位置
		fmt.Println(receipt.ContractAddress.Hex())          // 合约地址（为零地址表示非合约创建交易）
		break                                               // 仅处理第一个回执后退出循环
	}

	// 遍历通过哈希获取的回执列表（此处仅处理第一个回执后跳出循环）
	for _, receipt := range receiptsByNum {
		fmt.Println("交易状态：", receipt.Status)                // 交易状态：1 表示成功
		fmt.Println("日志事件:", receipt.Logs)                  // 交易触发的日志事件（空数组表示无事件）
		fmt.Println("交易哈希:", receipt.TxHash.Hex())          // 交易哈希的十六进制表示
		fmt.Println("在区块中的索引位置:", receipt.TransactionIndex) // 交易在区块中的索引位置
		fmt.Println(receipt.ContractAddress.Hex())          // 合约地址（为零地址表示非合约创建交易）
		break                                               // 仅处理第一个回执后退出循环
	}

	// 直接通过交易哈希获取特定交易的收据
	txHash := common.HexToHash(TX_HASH) // 交易hash
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
