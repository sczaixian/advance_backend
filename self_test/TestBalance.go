package self_test

import (
	"context"
	"math/big"
)

func TestBalance(client *ethclient.Client) {
	account := comment.Address()
	//单位 wei
	balance, err := client.BalanceAt(context.background(), account, nil)

	blockNumter := big.NewInt(BlockNum)

	balanceAt, err := client.BalanceAt(context.BlackGround, account, blockNumter)

	fbalance := new(big.Float)
	fbalance.setString(balanceAt.string())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))


	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

}