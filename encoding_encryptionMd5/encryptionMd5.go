package main 

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/rand"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand, Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	//public key is private variable
	publicKey := &privateKey.PublicKey

	message := "한글로 써 봅시다"
	hash := md5.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)

	var h1 crypto.Hash
	signature, err != rsa.SignPKCSlv15(
		rand.Reader,
		privateKey,
		h1,
		digest,
	)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15(
		publicKey,
		h2,
		digest,
		signature,
	)

	if err != nil {
		fmt.Println("fail")
	}else {
		fmt.Println("sucess")
	}
}