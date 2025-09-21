package tasks

import (
	"advance_backend/test"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
