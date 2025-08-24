package test

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestSubscribe() {
	// 1. 创建到以太坊测试网络（Ropsten）的WebSocket连接
	// WebSocket连接支持实时订阅功能，适合监听新区块等实时事件
	// 注意：Infura的WebSocket端点使用wss://协议
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err) // 连接失败时终止程序
	}

	// 2. 创建一个通道，用于接收新的区块头
	// 这个通道将传递*types.Header类型的指针
	headers := make(chan *types.Header)

	// 3. 订阅新的区块头事件
	// 当有新区块被挖出时，区块头信息将通过headers通道传递
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// 4. 无限循环，持续监听新区块事件
	for {
		// 使用select语句同时监听多个通道
		select {
		// 5. 处理订阅错误
		case err := <-sub.Err():
			log.Fatal(err) // 如果订阅出错，记录错误并退出

		// 6. 接收到新的区块头
		case header := <-headers:
			// 打印区块哈希的十六进制表示
			fmt.Println("区块哈希:", header.Hash().Hex())

			// 7. 通过区块哈希获取完整的区块信息
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			// 8. 打印区块详细信息
			fmt.Println("完整区块哈希:", block.Hash().Hex())      // 区块哈希
			fmt.Println("区块高度:", block.Number().Uint64())   // 区块号/高度
			fmt.Println("区块时间戳:", block.Time())             // 区块时间戳（Unix时间）
			fmt.Println("区块随机数:", block.Nonce())            // 工作量证明随机数
			fmt.Println("交易数量:", len(block.Transactions())) // 区块中包含的交易数量

			// 9. 可以在这里添加更多区块信息的处理逻辑
			// 例如：遍历交易、分析交易内容等
		}
	}

	// 注意：这个程序会无限运行，需要手动终止（Ctrl+C）
	// 在生产环境中，应该添加适当的退出机制
}
