//RSA加密和解密
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	//生成公钥
	publicKey := privateKey.PublicKey
	//根据公钥加密
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte("123"),//需要加密的字符串
		nil)
	if err != nil {
		panic(err)
	}


	fmt.Println("encrypted bytes: ", encryptedBytes)

	fmt.Println("encrypted strings 111 ", string(encryptedBytes))

	//根据私钥解密
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypted message: ", string(decryptedBytes))
}

