package xxx

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	//"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip39"
)

// 生成助记词
func GenerateMnemonic() (string, error) {
	// 生成128位熵(16字节)，对应12个单词的助记词
	//entropy, err := bip39.NewEntropy(128)
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

// 从助记词生成以太坊钱包
func EthWalletFromMnemonic(mnemonic string) error {
	// 验证助记词有效性
	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("无效的助记词")
	}

	// 从助记词生成种子
	seed := bip39.NewSeed(mnemonic, "") // 这里可以添加密码短语作为第二个参数

	// 从种子生成HD钱包的主密钥
	// 注意: 这里简化了流程，实际应用中应该使用完整的BIP32/BIP44派生路径
	// 以太坊的标准路径是 m/44'/60'/0'/0/0

	// 这里我们使用种子的前32字节作为私钥(简化实现，生产环境应使用完整HD派生)
	if len(seed) < 32 {
		return fmt.Errorf("种子长度不足")
	}

	privateKey, err := crypto.ToECDSA(seed[:32])
	if err != nil {
		return err
	}

	// 显示钱包信息
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("以太坊私钥：-->", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("无法转换公钥类型")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("以太坊公钥:", hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("以太坊地址:", address)

	return nil
}

// 从助记词生成Solana钱包
func SolanaWalletFromMnemonic(mnemonic string) error {
	// 验证助记词有效性
	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("无效的助记词")
	}

	// 从助记词生成种子
	seed := bip39.NewSeed(mnemonic, "") // 这里可以添加密码短语作为第二个参数

	// 使用种子的前32字节作为私钥(简化实现)
	// 注意: 实际Solana钱包使用不同的派生方法，这里仅作演示
	if len(seed) < 32 {
		return fmt.Errorf("种子长度不足")
	}

	//privateKey := seed[:32]

	privateKey := ed25519.NewKeyFromSeed(seed[:32])
	// 将私钥转换为base58编码的字符串
	base58PrivateKey := base58.Encode(privateKey)

	// 从私钥生成钱包
	//wallets, err := solana.WalletFromPrivateKey(privateKey)
	wallets, err := solana.WalletFromPrivateKeyBase58(base58PrivateKey)
	if err != nil {
		return err
	}

	fmt.Println("Solana私钥:", hexutil.Encode(privateKey))
	fmt.Println("Solana公钥:", wallets.PublicKey())

	return nil
}

func EthWalletCreateDeepseek() {
	// 生成助记词
	mnemonic, err := GenerateMnemonic()
	if err != nil {
		panic(err)
	}

	fmt.Println("生成的助记词:", mnemonic)
	fmt.Println("========== 以太坊钱包 ==========")

	// 从助记词生成以太坊钱包
	err = EthWalletFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}
}

func SolanaWalletCreateDeepseek() {
	// 生成助记词
	mnemonic, err := GenerateMnemonic()
	if err != nil {
		panic(err)
	}

	fmt.Println("生成的助记词:", mnemonic)
	fmt.Println("========== Solana钱包 ==========")

	// 从助记词生成Solana钱包
	err = SolanaWalletFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}

	//创建RPC客户端并请求空投(保持不变)
	account := solana.NewWallet()
	client := rpc.New(rpc.TestNet_RPC)

	out, err := client.RequestAirdrop(
		context.TODO(),
		account.PublicKey(),
		solana.LAMPORTS_PER_SOL*1,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		//panic(err)
		fmt.Println("空投err--SolanaWalletCreateDeepseek-->  :", err)
	}
	fmt.Println("空投交易签名:", out)
}
