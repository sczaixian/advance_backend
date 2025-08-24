package test

import (
	"context"      // 用于创建上下文，控制API调用超时和取消
	"crypto/ecdsa" // 椭圆曲线数字签名算法
	"fmt"
	"log"
	"math/big" // 大整数处理，用于处理以太坊的大数值

	"github.com/ethereum/go-ethereum/common"     // 以太坊常用类型（如地址、哈希）
	"github.com/ethereum/go-ethereum/core/types" // 核心类型（如交易）
	"github.com/ethereum/go-ethereum/crypto"     // 加密函数
	"github.com/ethereum/go-ethereum/ethclient"  // 以太坊客户端
)

func TestTransfer() {
	// 1. 连接到以太坊测试网络（Rinkeby）
	// Infura是一个提供以太坊节点服务的平台
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err) // 连接失败时终止程序
	}

	// 2. 从十六进制字符串加载私钥
	// 警告：在实际应用中，绝不应硬编码私钥，而应从安全存储中获取
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	// 3. 从私钥获取公钥
	publicKey := privateKey.Public()
	// 类型断言，确保公钥是ECDSA类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法断言类型: 公钥不是*ecdsa.PublicKey类型")
	}

	// 4. 从公钥生成以太坊地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 5. 获取账户的当前nonce值（交易计数器）
	// nonce是防止重放攻击的重要机制，每发送一笔交易nonce值增加1
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 6. 设置转账金额（1 ETH）
	// 以太坊使用wei作为最小单位，1 ETH = 10^18 wei
	value := big.NewInt(1000000000000000000) // 1 ETH (以wei为单位)

	// 7. 设置Gas限制（普通转账的标准Gas限制为21000）
	gasLimit := uint64(21000) // 单位

	// 8. 获取当前建议的Gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 9. 设置接收方地址
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")

	// 10. 创建交易数据（普通转账为空，合约调用需要包含调用数据）
	var data []byte

	txData := &types.LegacyTx{
		Nonce:    nonce,      // 从状态数据库获取当前 Nonce
		GasPrice: gasPrice,   // 20 Gwei big.NewInt(20000000000)
		Gas:      gasLimit,   // 标准转账 Gas
		To:       &toAddress, // 指向接收地址
		Value:    value,      // 0.5 ETH（单位：wei） big.NewInt(500000000000000000)
		Data:     data,       // 无附加数据
	}
	// 11. 创建交易对象
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTx(txData)

	// 12. 获取网络链ID（用于EIP-155签名，防止重放攻击跨链）
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 13. 使用私钥对交易进行签名
	// EIP-155签名包含链ID，确保交易只能在特定链上有效
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 14. 发送已签名的交易到以太坊网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// 15. 打印交易哈希（这是交易的唯一标识符）
	// 用户可以使用此哈希在区块链浏览器上查看交易状态
	fmt.Printf("交易已发送: %s", signedTx.Hash().Hex())
}
