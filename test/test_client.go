package test

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestClient() {
	// 使用ethclient包的Dial方法连接到指定的以太坊节点URL
	// 这里连接的是Cloudflare提供的公共以太坊节点
	client, err := ethclient.Dial("https://cloudflare-eth.com")

	// 检查连接是否出错
	if err != nil {
		// 如果连接出错，记录错误并终止程序
		log.Fatal(err)
	}

	// 连接成功时打印确认信息
	fmt.Println("we have a connection")

	// 使用空白标识符_忽略client变量（避免编译错误）
	// 在实际应用中，这里会使用client进行后续的区块链操作
	_ = client // 在后续章节中将会使用这个客户端
}
