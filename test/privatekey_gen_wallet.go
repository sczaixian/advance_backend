package test

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GenPrivateKey() {
	// 1. 生成一个新的ECDSA私钥（椭圆曲线数字签名算法）
	// 在以太坊中，私钥是一个256位（32字节）的随机数
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err) // 如果生成失败，记录错误并退出
	}

	// 2. 将私钥转换为字节切片，然后编码为十六进制字符串
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// hexutil.Encode会在结果前添加"0x"前缀，[2:]用于去掉这个前缀
	fmt.Println("私钥:", hexutil.Encode(privateKeyBytes)[2:])

	// 3. 从私钥获取对应的公钥
	publicKey := privateKey.Public()

	// 4. 将公钥转换为ECDSA类型（类型断言）
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法转换类型: 公钥不是*ecdsa.PublicKey类型")
	}

	// 5. 将ECDSA公钥转换为字节切片
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 以太坊公钥以0x04开头（表示未压缩格式），[4:]去掉这个前缀和第一个字节
	fmt.Println("从公钥字节生成的地址:", hexutil.Encode(publicKeyBytes)[4:])

	// 6. 使用内置函数直接从公钥生成以太坊地址
	// PubkeyToAddress函数内部实现了Keccak-256哈希和取后20字节的操作
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("通过PubkeyToAddress生成的地址:", address)

	// 7. 手动计算地址（验证地址生成过程）
	// 以太坊地址是公钥的Keccak-256哈希值的最后20个字节
	hash := sha3.NewLegacyKeccak256() // 创建Keccak-256哈希实例

	// 注意：这里跳过公钥的第一个字节（0x04），它是椭圆曲线点的前缀标识符
	// 在椭圆曲线密码学中，0x04表示未压缩的点
	hash.Write(publicKeyBytes[1:]) // 写入公钥数据（去掉前缀字节）

	// 计算哈希值
	hashBytes := hash.Sum(nil)
	fmt.Println("完整的Keccak-256哈希:", hexutil.Encode(hashBytes[:]))

	// 取哈希值的最后20个字节（40个十六进制字符）作为以太坊地址
	fmt.Println("手动计算的地址:", hexutil.Encode(hashBytes[12:]))
}
