package solana_and_eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"golang.org/x/crypto/sha3"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

//                区块链里其实是密钥管理器
//  是一个用于生成和管理『密钥对』、构建和签署『交易』，并与区块链网络进行交互的『工具』
//  私钥 →（生成）→ 公钥 →（生成）→ 地址
//  私钥 用于数字签名 当你想要发起一笔交易（比如转出比特币）时，你必须用私钥对这笔交易进行签名。
//  公钥 由私钥通过一种复杂的数学算法推导出来的一串字符。它可以公开
//       用来验证交易签名的有效性。网络上的其他人可以用你的公钥来验证你的交易签名是否确实由对应的私钥生成，从而确认交易合法。
//  地址 公钥再经过一系列哈希计算和编码后，生成的就是地址。它就像你的银行账号。
//  作用 公开给别人，用于接收加密货币。你可以放心地把你的地址告诉任何人，他们可以往这个地址转账，但没有私钥就无法转出。

//         “钱包”软件，其核心功能就是：
//  生成密钥对： 为你创建一套新的私钥、公钥和地址。
//  安全管理私钥： 帮你加密并安全地存储私钥（例如通过密码、助记词等方式）。
//  查询余额： 通过查询区块链网络，显示你所有地址上的资产总额。
//  构建和签署交易： 当你想要转账时，帮你构建交易信息，并用你的私钥进行签名。
//  广播交易： 将已签名的交易发送到区块链网络中，等待矿工（或验证者）打包确认

//           钱包分类
//  托管钱包： 交易所（如Binance, Coinbase）提供的钱包。他们替你保管私钥。你使用账户密码登录，交易由交易所执行。
//           这类似于银行，你信任的是交易所而不是自己。风险是如果交易所被黑客攻击或跑路，你的资产可能丢失。
//  自托管钱包： 你自己掌管私钥。这才是区块链意义上的真正钱包。上面提到的所有概念都适用于此。安全性责任完全在你自己身上。

//            助记词
//  钱包通常会使用一套标准算法，用12或24个常见的英文单词（称为助记词）来推导出你的私钥
//  助记词 = 私钥的另一种人类可读的形式

func EthWalletCreate() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("私钥：-->", hexutil.Encode(privateKeyBytes)[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address:", address)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("hash:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}

func SolanaWalletCreate() {
	account := solana.NewWallet()
	fmt.Println("account private key:", account.PrivateKey)
	fmt.Println("account public key:", account.PublicKey())

	// Create a new RPC client:
	client := rpc.New(rpc.DevNet_RPC)

	// Airdrop 1 SOL to the new account:
	out, err := client.RequestAirdrop(
		context.TODO(),
		account.PublicKey(),
		solana.LAMPORTS_PER_SOL*1,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		fmt.Println("error:--SolanaWalletCreate--> ", err)
		return
	}
	fmt.Println("airdrop transaction signature:", out)
}
