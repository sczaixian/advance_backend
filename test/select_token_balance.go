package test

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind" // 提供与合约交互的绑定功能
	"github.com/ethereum/go-ethereum/common"            // 以太坊常用类型（地址、哈希等）
	"github.com/ethereum/go-ethereum/ethclient"         // 以太坊客户端

	token "advance_backend/token" // 导入自动生成的ERC20代币合约绑定代码
	// 注意：这个路径需要根据实际情况调整，通常使用go generate生成
)

func SelectTokenBalance() {
	// 1. 连接到以太坊主网节点
	// 使用Cloudflare提供的公共以太坊节点
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err) // 连接失败时终止程序
	}

	// 2. 定义要查询的代币合约地址（这里使用Golem代币合约）
	tokenAddress := common.HexToAddress("0xfadea654ea83c00e5003d2ea15c59830b65471c0")

	// 3. 创建代币合约实例
	// 使用自动生成的合约绑定代码创建合约实例
	instance, err := token.NewToken(tokenAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	// 4. 定义要查询余额的以太坊地址
	address := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")

	// 5. 查询代币余额
	// &bind.CallOpts{} 提供调用选项（如区块高度，这里使用默认值即最新区块）
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	// 6. 查询代币名称
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// 7. 查询代币符号（缩写）
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// 8. 查询代币的小数位数
	// ERC20代币通常使用18位小数，与ETH相同
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// 9. 打印代币基本信息
	fmt.Printf("代币名称: %s\n", name)     // 输出: "代币名称: Golem Network"
	fmt.Printf("代币符号: %s\n", symbol)   // 输出: "代币符号: GNT"
	fmt.Printf("小数位数: %v\n", decimals) // 输出: "小数位数: 18"

	// 10. 打印原始余额（以最小单位表示）
	fmt.Printf("原始余额(wei): %s\n", bal) // 输出: "原始余额(wei): 74605500647408739782407023"

	// 11. 将余额从最小单位转换为可读格式
	fbal := new(big.Float)       // 创建大浮点数用于转换
	fbal.SetString(bal.String()) // 将余额的字符串表示设置为浮点数的值

	// 计算10^decimals的值（代币的精度基数）
	precision := math.Pow10(int(decimals))

	// 将余额除以精度基数，得到可读的余额值
	value := new(big.Float).Quo(fbal, big.NewFloat(precision))

	// 12. 打印格式化后的余额
	fmt.Printf("余额: %f", value) // 输出: "余额: 74605500.647409"
}
