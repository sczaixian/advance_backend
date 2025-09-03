package eth_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"
)

type MyData struct {
	Number uint
	Text   string
	Bytes  []byte
	BigInt *big.Int
}

func Test1() {
	// 编码一个结构体
	original := MyData{
		Number: 42,
		Text:   "Hello, RLP!",
		Bytes:  []byte{0x01, 0x02, 0x03},
		BigInt: big.NewInt(1234567890),
	}

	encodedBytes, err := rlp.EncodeToBytes(&original)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded: %x\n", encodedBytes)

	// 解码回结构体
	var decoded MyData
	err = rlp.DecodeBytes(encodedBytes, &decoded)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decoded Number: %d\n", decoded.Number)
	fmt.Printf("Decoded Text: %s\n", decoded.Text)
	fmt.Printf("Decoded Bytes: %x\n", decoded.Bytes)
	fmt.Printf("Decoded BigInt: %s\n", decoded.BigInt.String())
}
