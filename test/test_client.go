package test

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestClient() *ethclient.Client {
	// 使用ethclient包的Dial方法连接到指定的以太坊节点URL
	// 这里连接的是Cloudflare提供的公共以太坊节点
	client, err := ethclient.Dial(URL)

	// 检查连接是否出错
	if err != nil {
		// 如果连接出错，记录错误并终止程序
		log.Fatal(err)
	}

	// 连接成功时打印确认信息
	fmt.Println("we have a connection")

	// 使用空白标识符_忽略client变量（避免编译错误）
	// 在实际应用中，这里会使用client进行后续的区块链操作
	return client // 在后续章节中将会使用这个客户端
}

func TestClientLocal() *ethclient.Client {
	// npm install -g ganache-cli
	// ganache-cli
	// 使用相同的助记词 保证生成相同的 公开地址
	// ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func TestClientWebSocket() *ethclient.Client {
	// npm install -g ganache-cli
	// ganache-cli
	// 使用相同的助记词 保证生成相同的 公开地址
	// ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
	client, err := ethclient.Dial(WS_URL)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
