package test

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestBlock(client *ethclient.Client) {

	// 定义要查询的区块高度（9028872是Sepolia测试网上的一个具体区块）
	blockNumber := big.NewInt(BLOCK_HEIGHT)

	// 获取指定区块高度的区块头信息（Header只包含元数据，不包含交易详情）
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("获取区块头错误: ", err)
	}

	// 打印区块头信息：
	fmt.Println("区块高度:", header.Number.Uint64(), header.Number.String()) // 区块的编号 9028872
	fmt.Println("时间戳:", header.Time)                                      // 区块生成时的Unix时间戳  1755735216
	fmt.Println("难度值:", header.Difficulty.Uint64())                       // 该区块的工作量证明难度（PoS链上通常为0）  0
	fmt.Println("区块哈希:", header.Hash().Hex())                            // 区块的唯一哈希值标识 0xbe1af555f4e317adfb1d1894f2faed5c88777cd15aa6a195840a69198287f456

	// 获取完整的区块信息（包含所有交易详情）
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("获取完整区块错误: ", err)
	}

	// 打印完整区块信息（应与区块头数据一致）：
	fmt.Println("完整区块高度:", block.Number().Uint64())     // 确认区块高度  9028872
	fmt.Println("完整区块时间戳:", block.Time())              // 确认时间戳   1755735216
	fmt.Println("完整区块难度:", block.Difficulty().Uint64()) // 确认难度值     0
	// 0xbe1af555f4e317adfb1d1894f2faed5c88777cd15aa6a195840a69198287f456
	fmt.Println("完整区块哈希:", block.Hash().Hex())    // 确认哈希值（应与header.Hash()一致）
	fmt.Println("交易数量:", len(block.Transactions())) // 该区块中包含的交易数量  138

	// 通过交易哈希获取该区块的交易数量（另一种方法）
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal("获取交易数量错误: ", err)
	}
	fmt.Println("通过哈希验证的交易数量:", count) // 应与len(block.Transactions())结果一致  138

	// 注释中的哈希值可能是该区块中的某笔交易示例（实际代码中未使用）
	// 0x1fe7a460217e9525dd038d96c60b94c37e50bc6f2912e8b3ad27bc2671b6ec48
}
