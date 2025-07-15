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
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrEmbiMh6DoLNg3+xnAwvdQkXR
gZBjxgH2XfGhoxXw5w31bDVTUVqiyeHrn4O5tZQW+3X0dIk9crnkYhq41TkwxlZk
yn/Fn1YdBgx0Lv5AMjEJpBAgi+FX/pHznr2wM9UqAAO+vCh4rXVGwGCppMvg7uie
X6uf8VjTDIMwo8MmuwIDAQAB
-----END PUBLIC KEY-----`

	//const pubPEM = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrEmbiMh6DoLNg3+xnAwvdQkXRgZBjxgH2XfGhoxXw5w31bDVTUVqiyeHrn4O5tZQW+3X0dIk9crnkYhq41TkwxlZkyn/Fn1YdBgx0Lv5AMjEJpBAgi+FX/pHznr2wM9UqAAO+vCh4rXVGwGCppMvg7uieX6uf8VjTDIMwo8MmuwIDAQAB`

	msg := "{\"tagid\":\"1.305.29.1\",\"userid\":\"2c71488c-8c0c-4fcf-b369-0815841abd5f\",\"model\":\"SM-G965F\",\"country\":\"HK\"}"

	res, err := Encrypt(pubPEM, msg)

	if err != nil {
		fmt.Println("Encrypt error: err=%s", err.Error())
		return
	}

	fmt.Println("Encrypt: ", res)

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

	fmt.Println(3333, url.QueryEscape("hpEWNe2s6eaxjZkMI8aE3M1cVZ3+BFrX7IrowYWgEwYvauD9pVnZXW7zvvEc1bzTQSG/obu17tx3EOksdrO2YPEBmMspADrqUJVKW4gQy9t9PNXtycAKx+kVoXOAK6cVOwnzRL8daSTQCcGVd3I+vLJTxw4a6q9+xNf1lYBxoac="))

	return encodeString, nil
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