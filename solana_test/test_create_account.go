package solana_test

import (
	"context"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

/*
account private key: 4JRLKDf4XqHaqaap8rdXiC2rt6Ht2Lg9VDhnisTK7jd8dWMhnd2fYK3hJt5uy9cneh3U5ogEpsFAZXXAQtZSPw8z
account public key: E1iGi7Ww8SUAG92HN62rbJGEH7XQcGr6Hz4VcJ49ZeBG
airdrop transaction signature: 3pLpp47WiuoHPPPm6WFLNbhxme979ztduRM3VVSGxMcwv98p8xp15TP9nAATgcN9t9edCpC87Ue3Sxs9xBfxxY9J
*/

// TestCreateAccount 测试函数，用于演示如何在Solana区块链上创建新账户并请求空投
func TestCreateAccount() {
	// 创建一个新的Solana钱包账户
	// solana.NewWallet() 会生成一个新的随机密钥对（私钥和公钥）
	account := solana.NewWallet()

	// 打印新生成的私钥（Base58编码格式）
	// 警告：在实际生产环境中，永远不要公开或日志记录私钥
	fmt.Println("account private key:", account.PrivateKey)

	// 打印新生成的公钥（Base58编码格式）
	// 公钥是账户的地址，可以安全地公开分享
	fmt.Println("account public key:", account.PublicKey())

	// 创建与Solana区块链网络的RPC客户端连接
	// 使用开发网络(DevNet)而不是测试网络(TestNet)
	// DevNet是Solana的开发者网络，比TestNet更稳定，专门用于开发和测试
	client := rpc.New(rpc.DevNet_RPC)

	// 向新创建的账户请求空投1 SOL
	// RequestAirdrop是Solana RPC API的一个方法，用于在开发网络获取测试用的SOL
	// 参数说明：
	// - context.TODO(): 创建一个空的上下文，用于控制请求的取消和超时
	// - account.PublicKey(): 接收空投的目标账户公钥
	// - solana.LAMPORTS_PER_SOL*1: 空投金额，1 SOL = 10^9 lamports（lamports是Solana的最小单位）
	// - rpc.CommitmentFinalized: 指定交易确认级别为"finalized"，表示等待交易被网络最终确认
	out, err := client.RequestAirdrop(
		context.TODO(),
		account.PublicKey(),
		solana.LAMPORTS_PER_SOL*1,
		rpc.CommitmentFinalized,
	)

	// 错误处理：如果空投请求失败，抛出异常并终止程序
	if err != nil {
		panic(err)
	}

	// 打印空投交易的签名哈希
	// 这个签名可以用于在区块链浏览器上查询交易状态
	fmt.Println("airdrop transaction signature:", out)

	// 等待30秒，让空投交易有足够时间被网络确认和处理
	// 在实际应用中，应该使用更可靠的方法检查交易状态，而不是简单等待
	// 例如，可以使用client.GetSignatureStatuses()定期检查交易确认状态
	time.Sleep(30 * time.Second)

	// 可选：添加代码来验证空投是否成功
	// 可以查询账户余额确认是否收到了空投的SOL
}

func TestReceive1SOL() {
	client := rpc.New(rpc.DevNet_RPC)
	out, err := client.RequestAirdrop(
		context.TODO(),
		//byte[]("E1iGi7Ww8SUAG92HN62rbJGEH7XQcGr6Hz4VcJ49ZeBG"),
		byte("E1iGi7Ww8SUAG92HN62rbJGEH7XQcGr6Hz4VcJ49ZeBG"),
		solana.LAMPORTS_PER_SOL*1,
		rpc.CommitmentFinalized,
	)

	// 错误处理：如果空投请求失败，抛出异常并终止程序
	if err != nil {
		panic(err)
	}

	// 打印空投交易的签名哈希
	// 这个签名可以用于在区块链浏览器上查询交易状态
	fmt.Println("airdrop transaction signature:", out)

	// 等待30秒，让空投交易有足够时间被网络确认和处理
	// 在实际应用中，应该使用更可靠的方法检查交易状态，而不是简单等待
	// 例如，可以使用client.GetSignatureStatuses()定期检查交易确认状态
	time.Sleep(30 * time.Second)
}
