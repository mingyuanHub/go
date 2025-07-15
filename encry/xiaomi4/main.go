package main

import (
	"crypto/rsa"
	//"crypto/x509"
	//"encoding/pem"
	"fmt"
	"encoding/base64"
	"net/url"
	"math/big"
	"encoding/pem"
	"crypto/x509"
)

// NoPaddingRSAEncrypt 使用NoPadding方式进行RSA加密
func NoPaddingRSAEncrypt(pub *rsa.PublicKey, msg []byte) ([]byte, error) {
	// 确保消息长度合适
	if len(msg) > (pub.N.BitLen()+7)/8 {
		return nil, fmt.Errorf("message too long for RSA public key")
	}

	// 将消息转换为大整数
	m := new(big.Int).SetBytes(msg)

	// 执行RSA加密
	bigE := big.NewInt(int64(pub.E))
	c := new(big.Int).Exp(m, bigE, pub.N)

	// 将加密结果转换为字节切片
	return c.Bytes(), nil
}

const PUBLIC_KEY = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrEmbiMh6DoLNg3+xnAwvdQkXRgZBjxgH2XfGhoxXw5w31bDVTUVqiyeHrn4O5tZQW+3X0dIk9crnkYhq41TkwxlZkyn/Fn1YdBgx0Lv5AMjEJpBAgi+FX/pHznr2wM9UqAAO+vCh4rXVGwGCppMvg7uieX6uf8VjTDIMwo8MmuwIDAQAB
-----END PUBLIC KEY-----`

var txt = "{\"tagid\":\"1.305.29.1\",\"userid\":\"2c71488c-8c0c-4fcf-b369-0815841abd5f\",\"model\":\"SM-G965F\",\"country\":\"HK\"}"


func main() {
	block, _ := pem.Decode([]byte(PUBLIC_KEY))
	if block == nil || block.Type != "PUBLIC KEY" {
		fmt.Println(11111111)
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(222222)
	}

	// Assert that the public key is an *rsa.PublicKey
	rsaPubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		fmt.Println(3333)
	}

	// 要加密的消息
	msg := []byte(txt)

	// 使用NoPadding进行RSA加密
	encrypted, err := NoPaddingRSAEncrypt(rsaPubKey, msg)
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}

	fmt.Printf("Encrypted message: %x\n", encrypted)
	fmt.Println(base64.StdEncoding.EncodeToString(encrypted), err)
	fmt.Println(url.QueryEscape(base64.StdEncoding.EncodeToString(encrypted)), err)
}
