package test

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestAddress() {
	inputAddr := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"

	// 1. 验证地址格式
	if !common.IsHexAddress(inputAddr) {
		log.Fatal("Invalid Ethereum address format")
	}
	address := common.HexToAddress(inputAddr)
	// 直接调用 Hex() 方法输出 EIP-55 校验和地址
	fmt.Println("EIP-55 Address:", address.Hex()) // 输出：0x71C7656EC7ab88b098defB751B7401B5f6d8976F

	// 3. 计算正确哈希
	hash := crypto.Keccak256Hash(address.Bytes())
	fmt.Println("Keccak256 Hash:", hash.Hex())

	// 4. 按需输出字节数组
	fmt.Println("Bytes:", address.Bytes())
}

/*

client, err := ethclient.Dial("https://cloudflare-eth.com")
if err != nil {
	log.Fatal(err)
}

fmt.Println("we have a connection")
_ = client // we'll use this in the upcoming sections

address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
//fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
fmt.Println(address.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]

*/
