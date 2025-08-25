package self_test

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)


func GenPrivateKey(){
	privateKey, err := crypto.GenerateKey()


	privateKeyBytes := crypto.FromECDSA(privateKey)

	hexutil.Encode(privateKeyBytes)[2:]


	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	hexutil.Encode(publicKeyBytes)[4:]

	


}