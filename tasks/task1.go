package tasks

import (
	"advance_backend/test"
	"advance_backend/token"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
区块链读写 任务目标
使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。

	具体任务

环境搭建
安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
查询区块
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
输出查询结果到控制台。
发送交易
准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
对交易进行签名，并将签名后的交易发送到网络。
输出交易的哈希值。
*/
func Task1(blockNumber *big.Int) {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/860eeb1436364693ab37a4252f68f5ee")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("blockHash():", block.Hash())
	fmt.Println("block.Time():", block.Time())
	fmt.Println("transactions:", block.Transactions().Len())
	for _, tx := range block.Transactions() {
		fmt.Println("txHash():", tx.Hash().Hex())
		fmt.Println("tx.Gas():", tx.Gas())
		fmt.Println("tx.Data():", tx.Data())
		fmt.Println("tx.Value():", tx.Value())
		fmt.Println("tx.Nonce():", tx.Nonce())
		fmt.Println("tx.To():", tx.To())
		chainId := tx.ChainId()
		fmt.Println("chainId:", chainId)

		txx, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		if isPending {
			fmt.Println("交易仍在待处理池中（TxPool）", txx.To())
			continue
		}

		sender, err := client.TransactionSender(context.Background(), tx, block.Hash(), 0)
		if err != nil {
			signer := types.LatestSignerForChainID(chainId)
			sender, err = types.Sender(signer, tx)
			fmt.Println("sender1:", sender.Hex())
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("sender:", sender.Hex())
		break
	}

}

func Task1Tx() {
	//
	client, err := ethclient.Dial(test.URL)
	if err != nil {
		log.Fatal("连接失败: ", err)
	}

	privateKey, err := crypto.HexToECDSA("e039af6407e7622a8354bd45ea44de86ca663c81d6176ab698fe788e603b2682")
	if err != nil {
		log.Fatal("私钥格式错误: ", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(1000000000000000) // 0.001
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0xD930e6b5C1C1112d4c0Db00c4888557be1b58d0D")

	txData := &types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		GasPrice: gasPrice,
		Gas:      gasLimit,
	}
	tx := types.NewTx(txData)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signedTx:", signedTx.Hash().Hex())
}

/*
合约代码生成 任务目标
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。

	具体任务

编写智能合约
使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
编译智能合约，生成 ABI 和字节码文件。  solcjs --abi Counter.sol
使用 abigen 生成 Go 绑定代码
安装 abigen 工具。  go install github.com/ethereum/go-ethereum/cmd/abigen@latest     $env:Path += ";$env:GOPATH\bin"
使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。  abigen --abi=Counter_sol_Counter.abi --pkg=store --out=Counter.go
使用生成的 Go 绑定代码与合约交互
编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
调用合约的方法，例如增加计数器的值。
输出调用结果。
*/

const (
	contractAddr = "0x1307b1e71a527da6a990a141e4c382883470212a"
)

func CounterTest() {
	client, err := ethclient.Dial(test.URL)
	if err != nil {
		log.Fatal(err)
	}
	// 创建合约实例
	counterContract, err := token.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("6886c0e89cf2eca81e01cf1cb49462b6c18ac9d74992f77ce6192a9f87578835")
	if err != nil {
		log.Fatal("私钥格式错误: ", err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	opt.GasLimit = 3000000 // 足够的 gas

	// 调用合约方法
	tx, err := counterContract.IncrementCounter(opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent! Hash: %s\n", tx.Hash().Hex())

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("交易执行失败")
	}

	counter, err := counterContract.GetCounter(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Counter: %d\n", counter)

	if err != nil {
		log.Fatal(err)
	}

	tx, err = counterContract.DecrementedCounter(opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent! Hash: %s\n", tx.Hash().Hex())

	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("交易执行失败")
	}

	counter, err = counterContract.GetCounter(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Counter: %d\n", counter)
}

/*status	0x1 Transaction mined and execution succeed
transaction hash	0x21d3ca31273890f36ea1c23cb59bd3590228db42ea91efc29ba69cb68c68d270
block hash	0x34b419848c28c4c8db93ffe41e4e464eeb27db5697d7f3a3f6fa6f3439295186
block number	9252802
contract address	0x1307b1e71a527da6a990a141e4c382883470212a
from	0x3E0bDb54f94D735dDCf8D2074c852a8C22914aA7
to	Counter.(constructor)
gas	293615 gas
transaction cost	290145 gas
input	0x6080604052348015600e575f5ffd5b5061044f8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806312e968211461004e5780635b34b9661461005857806361bc221a146100625780638ada066e14610080575b5f5ffd5b61005661009e565b005b610060610131565b005b61006a610195565b6040516100779190610285565b60405180910390f35b61008861019a565b6040516100959190610285565b60405180910390f35b5f5f54116100e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100d8906102f8565b60405180910390fd5b5f5f8154809291906100f290610343565b91905055507fe3dbb27784766b8e76c133b768e79dce15035010e6f57b1da27f08c40879c6345f546040516101279190610285565b60405180910390a1565b620f42405f54106101455761014461022d565b5b5f5f8154809291906101569061036a565b91905055507f3cf8b50771c17d723f2cb711ca7dadde485b222e13c84ba0730a14093fad6d5c5f5460405161018b9190610285565b60405180910390a1565b5f5481565b5f620f42405f5411156101e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d9906103fb565b60405180910390fd5b5f5f541015610226576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021d906102f8565b60405180910390fd5b5f54905090565b5f5f819055507f4229cb99dd0cbc9a2bfac8f8f364d15c410e9b9541cdb6139b506fd657b925e55f546040516102639190610285565b60405180910390a1565b5f819050919050565b61027f8161026d565b82525050565b5f6020820190506102985f830184610276565b92915050565b5f82825260208201905092915050565b7f436f756e74657220697320746f6f20736d616c6c0000000000000000000000005f82015250565b5f6102e260148361029e565b91506102ed826102ae565b602082019050919050565b5f6020820190508181035f83015261030f816102d6565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61034d8261026d565b91505f820361035f5761035e610316565b5b600182039050919050565b5f6103748261026d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036103a6576103a5610316565b5b600182019050919050565b7f436f756e74657220697320746f6f206c617267650000000000000000000000005f82015250565b5f6103e560148361029e565b91506103f0826103b1565b602082019050919050565b5f6020820190508181035f830152610412816103d9565b905091905056fea26469706673582212208738cc0e5310e798d8aa0599eae9143ad32285ac2c794c772da032e16e16599464736f6c634300081e0033
decoded input	{}
decoded output	 -
logs	[]
raw logs	[]
view on Etherscan view on Blockscout*/

/*Transaction sent! Hash: 0xf07347b6bed770baf91115f9e0100eda4d78bcea316915d2a179573a894d658a
Counter: 1
2025/09/22 10:47:55 execution reverted*/
