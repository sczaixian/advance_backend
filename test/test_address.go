package test

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestAddress(client *ethclient.Client) common.Address {
	_ = client
	//inputAddr := ADDRESS_CMP_1
	inputAddr := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"

	// 1. 验证地址格式
	if !common.IsHexAddress(inputAddr) {
		log.Fatal("Invalid Ethereum address format")
	}
	address := common.HexToAddress(inputAddr)
	// 直接调用 Hex() 方法输出 EIP-55 校验和地址
	fmt.Println("EIP-55 Address:", address.Hex()) // 输出：0x71C7656EC7ab88b098defB751B7401B5f6d8976F

	// 3. 计算正确哈希
	hash_256 := crypto.Keccak256Hash(address.Bytes())
	fmt.Println("Keccak256Hash:", hash_256.Hex())

	// 替代：address.Hash().Hex() 因为 address没有 hash方法
	// 得到地址的哈希表示（32字节，前面填充0）
	hash := common.BytesToHash(address.Bytes())
	fmt.Println("BytesToHash:", hash.Hex())

	// 4. 按需输出字节数组
	fmt.Println("Bytes:", address.Bytes())

	return address
}
