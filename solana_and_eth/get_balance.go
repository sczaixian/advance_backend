package solana_and_eth

import (
	"advance_backend/test"
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func wei2eth(balance *big.Int) *big.Float {
	//fbalance := new(big.Float)
	//fbalance.SetString(balance.String())
	fbalance, _ := new(big.Float).SetString(balance.String())
	ethVal := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethVal
}

func lam2sol(balance uint64) *big.Float {
	lamports := new(big.Float).SetUint64(uint64(balance))
	solBalance := new(big.Float).Quo(lamports, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))
	return solBalance
}

func EthBalance(client *ethclient.Client) {
	account := common.HexToAddress(test.ADDRESS_CMP_1)
	balance, err := client.BalanceAt(context.Background(), account, nil) // 账户地址
	if err != nil {
		fmt.Println("err ----account_address---EthBalance------>> ", err)
	}
	fmt.Println("eth:", wei2eth(balance), ", wei:", balance)

	blockNumber := big.NewInt(5532993)
	balance, err = client.BalanceAt(context.Background(), account, blockNumber) // 区块链高度
	if err != nil {
		fmt.Println("err ----blockNumber---EthBalance------>> ", err)
	}
	fmt.Println("eth:", wei2eth(balance), ", wei:", balance)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		fmt.Println("err ----pendingBalance---EthBalance------>> ", err)
	}
	fmt.Println("eth:", wei2eth(pendingBalance), ", wei:", pendingBalance)
}

func SolanaBalance() {
	client := SolanaClientMainNet()
	pubKey := solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932")
	out, err := client.GetBalance(context.TODO(), pubKey, rpc.CommitmentFinalized)
	if err != nil {
		fmt.Println("err ----solanaBalance------>> ", err)
	}
	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports
	fmt.Println("◎", lam2sol(out.Value).Text('f', 10))
}
