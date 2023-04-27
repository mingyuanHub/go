package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"encoding/base64"
)

func main()  {
	buf :=
`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxppL+DjD4EvCp6XY805Q
SBhqN9bX8w12s9ctdkHEQKj4ufSwG1ffa6J79ip0w6sHNBTD2ZxaJFOTFGHugmJB
F0NTPEHixgDtWfUZxnpGhlU7u4sZPCzAOSjf1SeLmxRN6PhdFeAOsC8UbdZPe5vg
aGOT3GUWm8z3Lm7zjNayvIgsU/9rs//vRmZrn/6QSGoQUytfLGK+02hSxWBl/Um4
L/FJbE04Js0pdl90n4tkpHgfoo/uPAAoWmqBf2lz/TvEyKgZrGEn+6sxfzShIc/c
/pIKHfcP7rX7CzIq+p33E8zj+0U69wvswZY479guFN3VOz6l5iN+Wcu7IinuDQrV
cQIDAQAB
-----END PUBLIC KEY-----`
	//pem解码
	block, _ := pem.Decode([]byte(buf))
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	//类型断言
	publicKey:=publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte("333"))
	if err!=nil{
		panic(err)
	}

	encodeString := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Println(encodeString)
}
