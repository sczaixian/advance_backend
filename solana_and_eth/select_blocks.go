package solana_and_eth

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethRpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/gagliardetto/solana-go/rpc"
)

func EthSelectBlock(client *ethclient.Client) *types.Block {
	// HeaderByNumber
	header, err := client.HeaderByNumber(context.Background(), nil) // 传 nil 查最高的
	if err != nil {
		log.Fatal("Eth header: ", err)
	}

	//fmt.Println(header.Number.Int64(), header.ParentHash.Hex(), header.Hash().Hex(), header.Coinbase.Hex(),
	//	header.Root.Hex(), header.TxHash.Hex(), header.ReceiptHash.Hex(),
	//	header.Difficulty.String(), string(header.GasLimit), header.GasUsed,
	//	time.Unix(int64(header.Time), 0).Format("2006-01-02"),
	//	header.Extra, header.BaseFee.String(), header.Size().String(), header.MixDigest, header.Nonce)

	ot_str := fmt.Sprintf(
		"header:-> Block number: %d; "+
			"ParentHash：%s; Hash: %s\n"+
			"Coinbase: %s; "+
			"Root: %s; \n"+
			"TxHash: %s; "+
			"ReceiptHash: %s; \n"+
			"Difficulty: %s; "+
			"GasLimit: %s;\n "+
			"GasUsed: %d; "+
			"Time: %s; \n"+
			"Extra: %s; BaseFee: %s; \nSize: %s; "+
			"MixDigest: %s; \n"+
			"Nonce: %s;",
		header.Number.Int64(), header.ParentHash.Hex(), header.Hash().Hex(), header.Coinbase.Hex(),
		header.Root.Hex(), header.TxHash.Hex(), header.ReceiptHash.Hex(),
		header.Difficulty.String(), string(header.GasLimit), header.GasUsed,
		time.Unix(int64(header.Time), 0).Format("2006-01-02"),
		header.Extra, header.BaseFee.String(), header.Size().String(), header.MixDigest, header.Nonce)

	fmt.Println(ot_str)

	//-todo:----------通关 block hash    或者   block number  拿到 收据------------------------------
	//blockHash := common.HexToHash("块hash")
	blockNrOrHash := ethRpc.BlockNumberOrHashWithHash(header.Hash(), false)
	receiptsByHash, err := client.BlockReceipts(context.Background(), blockNrOrHash) //  得到 Receipt[]
	if err != nil {
		log.Fatal("receiptsByHash: ", err)
	}
	fmt.Println("blockNrOrHash.Hash(): ", blockNrOrHash.BlockHash.Hex(), blockNrOrHash.BlockNumber)

	blockNb := ethRpc.BlockNumber(header.Number.Int64() - 200)
	bnohwn := ethRpc.BlockNumberOrHashWithNumber(blockNb)

	//fmt.Println("bnohwn.Hash(): ", bnohwn.BlockHash.Hex())
	receiptsByNum, err := client.BlockReceipts(context.Background(), bnohwn) //  得到 Receipt[]

	fmt.Println("---receiptsByHash[0] == receiptsByNum[0]--->> ", receiptsByHash[0] == receiptsByNum[0])

	//block, err := client.BlockByNumber(context.Background(), big.NewInt(header.Number.Int64()))
	block, err := client.BlockByNumber(context.Background(), big.NewInt(2394201))
	if err != nil {
		log.Fatal("BlockByNumber:", err)
	}

	block_str := fmt.Sprintf("block: transactions:%d; Hash: %s; \n"+
		"txhash: %s; nonce: %d; \n"+
		"number: %s; difficulty: %d; \n"+
		"gaslimit: %d; gasUsed:%d; \n time:%s; ",
		block.Transactions().Len(),
		block.Hash().String(), block.TxHash().String(), block.Nonce(),
		block.Number().String(), block.Difficulty().String(),
		block.GasLimit(), block.GasUsed(), time.Unix(int64(block.Time()), 0).Format("2006-01-02"))

	fmt.Println(block_str)
	fmt.Println(len(block.Transactions())) // 70

	return block
}

func SolanaSelectBlock(client *rpc.Client) *rpc.GetBlockResult {
	example, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		log.Fatal("GetLatestBlockhash: ", err)
	}
	//out, err := client.GetBlock(context.TODO(), uint64(example.Context.Slot))
	out, err := client.GetBlockWithOpts(context.TODO(), uint64(example.Context.Slot), &rpc.GetBlockOpts{
		MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion0, // 版本0
	})

	if err != nil {
		log.Fatal("SolanaGetBlock: ", err)
	}
	//fmt.Println(out.Blockhash.String(), out.PreviousBlockhash.String(), out.ParentSlot, len(out.Transactions),
	//	len(out.Signatures), len(out.Rewards), time.Unix(int64(*out.BlockTime), 0).Format("2006-01-02"),
	//	out.BlockHeight, out.NumRewardPartitions)
	ot_str := fmt.Sprintf(
		"Blockhash:%s; PreviousBlockhash:%s; \n"+
			" ParentSlot:%d; TransactionsLen:%d; \n"+
			"SignaturesLen:%d; Rewards[BlockReward](len):%d; \n"+
			"BlockTime:%s; BlockHeight:%d; \n NumRewardPartitions:%d; \n",
		out.Blockhash.String(), out.PreviousBlockhash.String(), out.ParentSlot, len(out.Transactions),
		len(out.Signatures), len(out.Rewards), time.Unix(int64(*out.BlockTime), 0).Format("2006-01-02"),
		out.BlockHeight, out.NumRewardPartitions)

	fmt.Println(ot_str)
	return out
}
