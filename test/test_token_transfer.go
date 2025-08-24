package test

import (
	"context"      // 用于创建上下文，控制API调用超时和取消
	"crypto/ecdsa" // 椭圆曲线数字签名算法
	"fmt"
	"log"
	"math/big" // 大整数处理，用于处理以太坊的大数值

	"golang.org/x/crypto/sha3" // SHA3哈希函数（Keccak-256）

	"github.com/ethereum/go-ethereum"                // 以太坊核心接口
	"github.com/ethereum/go-ethereum/common"         // 以太坊常用类型（如地址、哈希）
	"github.com/ethereum/go-ethereum/common/hexutil" // 十六进制编码工具
	"github.com/ethereum/go-ethereum/core/types"     // 核心类型（如交易）
	"github.com/ethereum/go-ethereum/crypto"         // 加密函数
	"github.com/ethereum/go-ethereum/ethclient"      // 以太坊客户端
)

func TestTokenTransfer() {
	// 1. 连接到以太坊Sepolia测试网络
	// 使用Alchemy提供的节点服务（需要替换为有效的API密钥）
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err)
	}

	// 2. 从十六进制字符串加载私钥
	// 警告：在实际应用中，绝不应硬编码私钥，而应从安全存储中获取
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY_HOM1) // 需要替换为实际私钥
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

	// 4. 从公钥生成以太坊地址（发送方地址）
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 5. 获取账户的当前nonce值（交易计数器）
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 6. 设置转账金额（0 ETH，因为这是代币转账，不是ETH转账）
	value := big.NewInt(0) // 0 wei (0 ETH)

	// 7. 获取当前建议的Gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 8. 设置代币接收方地址
	toAddress := common.HexToAddress(ADDRESS_HOM_1)

	// 9. 设置代币合约地址
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	// 10. 构建ERC20代币转账的函数调用数据
	// 首先获取transfer函数的选择器（前4个字节）
	transferFnSignature := []byte("transfer(address,uint256)") // 函数签名
	hash := sha3.NewLegacyKeccak256()                          // 创建Keccak-256哈希实例
	hash.Write(transferFnSignature)                            // 写入函数签名
	methodID := hash.Sum(nil)[:4]                              // 取前4个字节作为方法ID
	fmt.Println("方法ID:", hexutil.Encode(methodID))             // 输出: 0xa9059cbb

	// 11. 准备接收方地址参数（左填充到32字节）
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println("填充后的地址:", hexutil.Encode(paddedAddress)) // 输出: 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	// 12. 准备转账金额参数（左填充到32字节）
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000个代币（假设代币有18位小数）
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println("填充后的金额:", hexutil.Encode(paddedAmount)) // 输出: 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// 13. 组合所有数据：方法ID + 地址参数 + 金额参数
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 14. 估算交易所需的Gas限制
	// 注意：这里发送到代币合约地址，而不是普通接收方地址
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,   // 发送方地址
		To:   &tokenAddress, // 代币合约地址（注意是指针）
		Data: data,          // 调用数据
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("估算的Gas限制:", gasLimit) // 输出: 23256

	// 15. 创建交易对象
	// 注意：接收方是代币合约地址，而不是普通接收方
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	// 16. 获取网络链ID（用于EIP-155签名）
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 17. 使用私钥对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 18. 发送已签名的交易到以太坊网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// 19. 打印交易哈希
	fmt.Printf("交易已发送: %s", signedTx.Hash().Hex())
}
