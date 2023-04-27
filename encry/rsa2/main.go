package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/url"
)

func main()  {

	const pubPEM = `
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAK3qoHHDhWwukjwu9y59qTnqLJfurLm0
lhaw2aU4V2kIE3dBdL55Ld1SZwEZPB7t409X4tzd+vkVOy9jVOBmOCECAwEAAQ==
-----END PUBLIC KEY-----`

	const privatePEM = `
-----BEGIN PRIVATE KEY-----
MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAreqgccOFbC6SPC73
Ln2pOeosl+6subSWFrDZpThXaQgTd0F0vnkt3VJnARk8Hu3jT1fi3N36+RU7L2NU
4GY4IQIDAQABAkA6z8bl8apixPTqqS8pZ5EcZpYh4rJCMlE25yMSfhUBDQ1f6OnH
aHgijOG6JWtnqvY9q2+bJYakhbW/U3Hj8LZhAiEA1al0i/IRi0QT9jXbdYkZ7vUr
waZrW6R6BaYJTxaEs20CIQDQYPvbgsA3QIqPfENq/nSWhna8hli6D0Vy+j2F5lAz
BQIhAJn9sdQTOXXIMSLomi1SDPDenxTI3uOD3bYoftkTf7zZAiBCckqdiqoEdF46
tuNAoPdIcIQ4RZbRbcE1kro/kluMiQIgKnjnPAr4ySoCHqpSjuEqh4BFKq98670z
J/ZrICNdaT8=
-----END PRIVATE KEY-----`

	msg := "hahahahaa"

	res, err := Encrypt(pubPEM, msg)

	if err != nil {
		fmt.Println("Encrypt error: err=%s", err.Error())
		return
	}

	fmt.Println("Encrypt: ", res)

	msg1, err := Decrypt(privatePEM, res)

	if err != nil {
		fmt.Println("Decrypt error: err=%s", err.Error())
		return
	}

	fmt.Println("Decrypt: ", msg1)

}

func Encrypt(pubPEM, msg string) (string, error) {
	cipherByte, err := RsaEncrypt(pubPEM, msg)
	if err != nil {
		return "", err
	}

	fmt.Println("00000000", cipherByte)

	encodeString := base64.StdEncoding.EncodeToString(cipherByte)

	fmt.Println(111111111, encodeString)

	encodeString = url.QueryEscape(encodeString)

	fmt.Println(222222222, encodeString)

	return encodeString, nil
}

func Decrypt(privatePEM, msg string) (string, error) {

	msg, err := url.QueryUnescape(msg)
	if err != nil {
		return "", err
	}

	fmt.Println(3333333, msg)

	decodeString, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}

	fmt.Println(444444444, decodeString)


	cipherByte, err := RSADecrypt(privatePEM, decodeString)
	if err != nil {
		return "", err
	}

	return string(cipherByte), nil
}

func RsaEncrypt(pubPEM, msg string) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode([]byte(pubPEM))

	input := block.Bytes

	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(input)
	if err!=nil{
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(msg))
	if err!=nil{
		return nil, err
	}

	return cipherText, nil
}

func RSADecrypt(privatePEM string, msgEnc []byte) ([]byte, error){

	//pem解码
	block, _ := pem.Decode([]byte(privatePEM))

	//X509解码
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err!=nil{
		return nil, err
	}

	//对密文进行解密
	plainText, _:= rsa.DecryptPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), msgEnc)

	//返回明文
	return plainText, nil
}