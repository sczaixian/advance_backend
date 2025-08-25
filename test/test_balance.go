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

func TestBalance(client *ethclient.Client, account *common.Address) {

	// 获取账户当前最新区块的余额
	// context.Background() 提供空的上下文，不设置超时或取消
	// 第三个参数为nil表示获取最新区块的余额
	balance, err := client.BalanceAt(context.Background(), *account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("最新区块余额：", balance, "(wei)") // 输出以大整数形式表示的余额（单位: wei）

	// 获取账户在特定区块高度时的余额
	// 创建一个表示区块5532993的大整数
	blockNumber := big.NewInt(BLOCK_HEIGHT)
	// 查询该账户在指定区块高度时的余额
	balanceAt, err := client.BalanceAt(context.Background(), *account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	// 输出: 25729324269165216042 (单位: wei)
	fmt.Println("区块：", BLOCK_HEIGHT, "的余额：", balanceAt, "(wei)")

	// 将余额从wei转换为ETH
	// 1 ETH = 10^18 wei
	fbalance := new(big.Float)             // 创建一个大浮点数用于转换
	fbalance.SetString(balanceAt.String()) // 将余额的字符串表示设置为浮点数的值
	// 除以10^18将wei转换为ETH
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("转换后的余额：", ethValue, "(ETH)", ethValue.String()) // 输出: 25.729324269165216041 (单位: ETH)

	// 获取账户的待处理余额（包括内存池中尚未确认的交易）
	pendingBalance, err := client.PendingBalanceAt(context.Background(), *account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("待处理余额：", pendingBalance, "(wei)") // 输出: 25729324269165216042 (单位: wei)
}
