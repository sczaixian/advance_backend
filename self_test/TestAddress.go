

package self_test


func TestAddress(client * ethclient.Client){
	_ = client

	inputAddress := "xxx"
	if !common.IsHexAdderss(inputAddress) {
		log.Fatal("xxx")
	}

	address := comment.HexToAddress(inputAddress)
	// 输出校验和地址
	fmt.Println(address.Hex())

	hash_256 := crpyto.Keccak256Hash(address.Bytes())
	fmt.Println(hash_256.Hex())

	//填充后的地址
	hash := comment.BytesToHash(address.Bytes())
	fmt.Println(hash.Hex())

	return address
}