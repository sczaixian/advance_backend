package test

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestBalance() {
	// 创建以太坊客户端连接
	// 使用Cloudflare提供的公共以太坊节点
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	// 将十六进制地址字符串转换为Address类型
	// 这是一个以太坊账户地址
	account := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")

	// 获取账户当前最新区块的余额
	// context.Background() 提供空的上下文，不设置超时或取消
	// 第三个参数为nil表示获取最新区块的余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 输出以大整数形式表示的余额（单位: wei）

	// 获取账户在特定区块高度时的余额
	// 创建一个表示区块5532993的大整数
	blockNumber := big.NewInt(5532993)
	// 查询该账户在指定区块高度时的余额
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 输出: 25729324269165216042 (单位: wei)

	// 将余额从wei转换为ETH
	// 1 ETH = 10^18 wei
	fbalance := new(big.Float)             // 创建一个大浮点数用于转换
	fbalance.SetString(balanceAt.String()) // 将余额的字符串表示设置为浮点数的值
	// 除以10^18将wei转换为ETH
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 输出: 25.729324269165216041 (单位: ETH)

	// 获取账户的待处理余额（包括内存池中尚未确认的交易）
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pendingBalance) // 输出: 25729324269165216042 (单位: wei)
}
