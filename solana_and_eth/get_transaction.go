package solana_and_eth

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gagliardetto/solana-go/rpc"
)

func EthTransaction(client *ethclient.Client, block *types.Block) {
	for _, tx := range block.Transactions() {

		tx_str := fmt.Sprintf("Type: %s; ChainId: %d; \n"+
			"AccessList_len: %d; Data: %s; \n"+
			"Gas: %d; GasPrice: %d; \n"+
			"GasFeeCap: %d; GasTipCap: %d; \n"+
			"Value: %d; To: %s; \n"+
			"Nonce: %d; Hash: %s; \n"+
			"Size: %d;",
			tx.Type(), tx.ChainId(),
			len(tx.AccessList()), tx.Data(),
			tx.Gas(), tx.GasPrice(),
			tx.GasFeeCap(), tx.GasTipCap(),
			tx.Value(), tx.To().Hex(), tx.Nonce(),
			tx.Hash().Hex(), tx.Size())

		// 使用 EIP155Signer 还原出 sender 地址
		if sender, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx); err == nil {
			tx_str += fmt.Sprintf("\nSender: %s", sender.Hex())
		}

		if receipt, err := client.TransactionReceipt(context.Background(), tx.Hash()); err == nil {
			tx_str += fmt.Sprintf("\nReceipt: %d", receipt.Status) //  收据
		}
		fmt.Println(tx_str)
		break
	}
	//blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	// 块哈希和块内事务的索引值
	//count, err := client.TransactionCount(context.Background(), blockHash)
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), block.Hash(), idx)

		if err == nil {
			fmt.Println(tx.Hash().Hex())
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err == nil {
				if len(receipt.Logs) == 0 {
					continue
				}
				fmt.Println("----------------- receipt -------------------------")
				spew.Dump(receipt)
			}
		}

		break
	}
	// 在给定具体事务哈希值的情况下直接查询单个事务
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	// 交易信息， 交易状态师傅处在等待， 错误信息
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)
}

func SolanaTransaction(client *rpc.Client, block *rpc.GetBlockResult) {
	// EKchX9dFPQCuJRyLLytUbXURoXzMSH74ry55yjioWKdn
	// MustSignatureFromBase58, 出现问题 直接panic；  SignatureFromBase58 发生问题 返回err
	// 接收一个 Base58 编码的字符串 返回对应一个 64 字节的 Signature 对象
	// 你需要自行对返回的 Base64 字符串进行解码（使用 base64.StdEncoding.DecodeString），
	// 然后才能通过 solana.TransactionFromDecoder 或其他方法解析为可操作的交易对象。
	// 对人类不直接可读，需要额外解析。
	// 	给你更多的控制权。如果你只需要交易的原始字节数据（例如为了计算哈希、验证签名或进行自定义解析），获取 Base64 编码的数据可能更高效，避免了 RPC 节点进行解析的开销
	//txsig := solana.MustSignatureFromBase58("4bjVLV1g9SAfv7BSAdNnuSPRbSscADHFe4HegL6YVcuEBMY83edLEvtfjE4jfr6rwdLwKBQbaFiGgoLGtVicDzHq")
	//fmt.Println("-------  txsig  ----------->>", txsig)
	//out, err := client.GetTransaction(context.TODO(), txsig, &rpc.GetTransactionOpts{Encoding: solana.EncodingBase64})
	//if err == nil {
	//	spew.Dump(out)
	//}
	//decodedTx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(out.Transaction.GetBinary()))
	//if err == nil {
	//	spew.Dump(decodedTx)
	//}
	//// 返回一个已解析的 JSON 对象
	//out, err = client.GetTransaction(context.TODO(), txsig, nil)
	//if err == nil {
	//	spew.Dump(out)
	//}
	for _, tx := range block.Transactions {
		spew.Dump(tx)
		break
	}

}

/*
+-----------------------+--------------------------------------------+----------------------------------------------------+
| 特性                  | 以太坊 (Ethereum)                          | Solana                                             |
+-----------------------+--------------------------------------------+----------------------------------------------------+
| 核心概念              | 交易收据 (Transaction Receipt)             | 交易详情与元数据 (Transaction Details & Metadata) |
+-----------------------+--------------------------------------------+----------------------------------------------------+
| 获取方式              | eth_getTransactionReceipt                  | getTransaction （需查询）                         |
+-----------------------+--------------------------------------------+----------------------------------------------------+
| 状态确认               | 收据中的 status 字段                       | getSignatureStatuses 或                           |
|                       |                                            | getTransaction 返回的 meta.err 字段               |
+-----------------------+--------------------------------------------+----------------------------------------------------+
| Gas/费用消耗          | 收据中的 gasUsed                           | getTransaction 返回的 meta.fee 字段               |
+-----------------------+--------------------------------------------+----------------------------------------------------+
| 日志/事件             | 收据中的结构化 logs（事件）                 | getTransaction 返回的 meta.logMessages（字符串）  |
|                       |                                            | 或新版本交易中的结构化事件                         |
+-----------------------+--------------------------------------------+----------------------------------------------------+
| 余额变化              | 需要自行解析日志事件计算                    | getTransaction 返回的 meta.preBalances 和         |
|                       |                                            | meta.postBalances 清晰列出                         |
+-----------------------+--------------------------------------------+----------------------------------------------------+
*/

/*  以太坊的一笔交易
(*types.Receipt)(0xc0002243c0)({
    Type: (uint8) 2,
    PostState: ([]uint8) <nil>,
    Status: (uint64) 1,
    CumulativeGasUsed: (uint64) 273848,
    Bloom: (types.Bloom) (len=256 cap=256) {
        00000000  00 00 00 00 00 20 00 00  80 20 00 00 00 00 00 00  |..... ... ......|
        00000010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000020  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000030  00 40 20 00 00 00 00 00  00 00 00 00 00 20 00 00  |.@ .......... ..|
        00000040  00 04 00 00 00 00 00 00  00 00 00 08 00 08 00 00  |................|
        00000050  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000060  00 00 00 00 02 00 00 00  00 00 00 00 00 00 08 00  |................|
        00000070  00 00 00 00 00 00 18 00  00 00 00 10 00 00 00 00  |................|
        00000080  00 00 00 02 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000090  00 00 00 00 00 00 00 00  00 00 00 00 08 00 00 00  |................|
        000000a0  02 00 00 00 00 00 01 00  00 00 04 00 00 00 00 00  |................|
        000000b0  80 00 00 00 10 00 00 00  00 00 00 00 00 00 00 00  |................|
        000000c0  00 00 00 02 00 01 00 00  00 00 00 00 00 04 00 00  |................|
        000000d0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 20 00  |.............. .|
        000000e0  00 10 00 00 00 00 00 00  00 00 00 00 00 00 00 02  |................|
        000000f0  00 00 40 00 00 00 00 00  00 00 00 00 00 01 00 00  |..@.............|
    },
    Logs: ([]*types.Log) (len=3 cap=4) {
        (*types.Log)(0xc0001a40b0)({
            Address: (common.Address) (len=20 cap=20) 0x969D499507B4f437953Db24A4980FdEEDa6Db8a1,
            Topics: ([]common.Hash) (len=3 cap=4) {
                (common.Hash) (len=32 cap=32) 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925,
                (common.Hash) (len=32 cap=32) 0x000000000000000000000000d526c2ad01421951aad98500ab413eaa6aba54af,
                (common.Hash) (len=32 cap=32) 0x00000000000000000000000089a5f2c62213b18ee5f83b21cb1a323920c9b101
            },
            Data: ([]uint8) (len=32 cap=32) {
                00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
                00000010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
            },
            BlockNumber: (uint64) 2394201,
            TxHash: (common.Hash) (len=32 cap=32) 0xe44e5606488c5656f63a982b134825fad4d063fe8fffaba272472873817408c4,
            TxIndex: (uint) 0,
            BlockHash: (common.Hash) (len=32 cap=32) 0xfd93dcd1233ba2d24bf34b8712b9479a3fd7c7fc29d5b90ee42aaa9b10e7cace,
            BlockTimestamp: (uint64) 1669912896,
            Index: (uint) 0,
            Removed: (bool) false
        }),
        (*types.Log)(0xc0001a4160)({
            Address: (common.Address) (len=20 cap=20) 0x969D499507B4f437953Db24A4980FdEEDa6Db8a1,
            Topics: ([]common.Hash) (len=3 cap=4) {
                (common.Hash) (len=32 cap=32) 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef,
                (common.Hash) (len=32 cap=32) 0x000000000000000000000000d526c2ad01421951aad98500ab413eaa6aba54af,
                (common.Hash) (len=32 cap=32) 0x00000000000000000000000089a5f2c62213b18ee5f83b21cb1a323920c9b101
            },
            Data: ([]uint8) (len=32 cap=32) {
                00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
                00000010  00 00 00 00 00 00 00 00  45 63 91 82 44 f4 00 00  |........Ec..D...|
            },
            BlockNumber: (uint64) 2394201,
            TxHash: (common.Hash) (len=32 cap=32) 0xe44e5606488c5656f63a982b134825fad4d063fe8fffaba272472873817408c4,
            TxIndex: (uint) 0,
            BlockHash: (common.Hash) (len=32 cap=32) 0xfd93dcd1233ba2d24bf34b8712b9479a3fd7c7fc29d5b90ee42aaa9b10e7cace,
            BlockTimestamp: (uint64) 1669912896,
            Index: (uint) 1,
            Removed: (bool) false
        }),
        (*types.Log)(0xc0001a4210)({
            Address: (common.Address) (len=20 cap=20) 0xe8B0a865E4663636BF4D6b159c57333210b0C229,
            Topics: ([]common.Hash) (len=4 cap=4) {
                (common.Hash) (len=32 cap=32) 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8,
                (common.Hash) (len=32 cap=32) 0x0000000000000000000000000000000000000000000000000000000000000000,
                (common.Hash) (len=32 cap=32) 0x0000000000000000000000969d499507b4f437953db24a4980fdeeda6db8a102,
                (common.Hash) (len=32 cap=32) 0x000000000000000000000000000000000000000000000000000000000000122b
            },
            Data: ([]uint8) {
            },
            BlockNumber: (uint64) 2394201,
            TxHash: (common.Hash) (len=32 cap=32) 0xe44e5606488c5656f63a982b134825fad4d063fe8fffaba272472873817408c4,
            TxIndex: (uint) 0,
            BlockHash: (common.Hash) (len=32 cap=32) 0xfd93dcd1233ba2d24bf34b8712b9479a3fd7c7fc29d5b90ee42aaa9b10e7cace,
            BlockTimestamp: (uint64) 1669912896,
            Index: (uint) 2,
            Removed: (bool) false
        })
    },
    TxHash: (common.Hash) (len=32 cap=32) 0xe44e5606488c5656f63a982b134825fad4d063fe8fffaba272472873817408c4,
    ContractAddress: (common.Address) (len=20 cap=20) 0x0000000000000000000000000000000000000000,
    GasUsed: (uint64) 273848,
    EffectiveGasPrice: (*big.Int)(0xc00089e0a0)(1500000007),
    BlobGasUsed: (uint64) 0,
    BlobGasPrice: (*big.Int)(<nil>),
    BlockHash: (common.Hash) (len=32 cap=32) 0xfd93dcd1233ba2d24bf34b8712b9479a3fd7c7fc29d5b90ee42aaa9b10e7cace,
    BlockNumber: (*big.Int)(0xc00089e060)(2394201),
    TransactionIndex: (uint) 0
})
*/
/*   solana的一笔交易
(rpc.TransactionWithMeta) {
    Slot: (uint64) 0,
    BlockTime: (*solana.UnixTimeSeconds)(<nil>),
    Transaction: (*rpc.DataBytesOrJSON)(0xc000340190)({
        rawDataEncoding: (solana.EncodingType) (len=6) "base64",
        asDecodedBinary: (solana.Data) AbrTs5deHH8187WOjZrVuDdCCykKCRWt7RNcZV3lf6NVLkIQA+QlIwFA65qxI/ofDif5f+x72iYDebuKKsqhPQABAAED/ailItoEdsrkNQt8xiVPgRL/WC23yWcF1ztdUU0H8THHqzcmZ/iTQa0r5+UBLhVQBm3W9SLgn8Ds/GTlqwphpgdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAAiFecgfk0uoL1Q14859XZNkgMZMUQMX9AnY2couQ0xkBAgIBAJQBDgAAAK5FKxUAAAAAHwEfAR4BHQEcARsBGgEZARgBFwEWARUBFAETARIBEQcQAQ8BDgENAQwBCwEKAQkBCAEHAQYBBQEEAQMBAgEBPy8Hr0Leame+Lyu9FzaSM14EEcQ9jAnAcyFZHE1XFo8Bdpe2aAAAAADxFe5axI812hDJLi6Ph/HF075by/RJNoGpHpm0ZiZ7g==,
        asJSON: (json.RawMessage) <nil>
    }),
    Meta: (*rpc.TransactionMeta)(0xc00013a000)({
        Err: (interface {}) <nil>,
        Fee: (uint64) 5000,
        PreBalances: ([]uint64) (len=3 cap=4) {
            (uint64) 17649381532987,
            (uint64) 18026858640,
            (uint64) 1
        },
        PostBalances: ([]uint64) (len=3 cap=4) {
            (uint64) 17649381527987,
            (uint64) 18026858640,
            (uint64) 1
        },
        InnerInstructions: ([]rpc.InnerInstruction) {
        },
        PreTokenBalances: ([]rpc.TokenBalance) {
        },
        PostTokenBalances: ([]rpc.TokenBalance) {
        },
        LogMessages: ([]string) (len=2 cap=2) {
            (string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
            (string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
        },
        Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
            (string) (len=2) "Ok": (interface {}) <nil>
        },
        Rewards: ([]rpc.BlockReward) {
        },
        LoadedAddresses: (rpc.LoadedAddresses) {
            ReadOnly: (solana.PublicKeySlice) {
            },
            Writable: (solana.PublicKeySlice) {
            }
        },
        ReturnData: (rpc.ReturnData) {
            ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
            Data: (solana.Data)
        },
        ComputeUnitsConsumed: (*uint64)(0xc00040e710)(2100)
    }),
    Version: (rpc.TransactionVersion) -1
}

*/
