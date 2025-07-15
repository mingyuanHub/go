package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/sha256"
	"encoding/pem"
	"encoding/base64"
	"fmt"
	"net/url"
)

type noPadding struct{}

func (n noPadding) Pad(buf []byte) []byte {
	return buf
}

func (n noPadding) Unpad(buf []byte) ([]byte, error) {
	return buf, nil
}

// 使用rsa.Encrypt进行加密
func rsaEncryptNoPadding(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	padding := noPadding{}
	msgBlock := padding.Pad(msg)
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, msgBlock, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// 使用rsa.Decrypt进行解密
func rsaDecryptNoPadding(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	padding := noPadding{}
	msgBlock, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return padding.Unpad(msgBlock)
}

const PUBLIC_KEY = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrEmbiMh6DoLNg3+xnAwvdQkXR
gZBjxgH2XfGhoxXw5w31bDVTUVqiyeHrn4O5tZQW+3X0dIk9crnkYhq41TkwxlZk
yn/Fn1YdBgx0Lv5AMjEJpBAgi+FX/pHznr2wM9UqAAO+vCh4rXVGwGCppMvg7uie
X6uf8VjTDIMwo8MmuwIDAQAB
-----END PUBLIC KEY-----`

const PUBLIC_KEY2 = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrEmbiMh6DoLNg3+xnAwvdQkXRgZBjxgH2XfGhoxXw5w31bDVTUVqiyeHrn4O5tZQW+3X0dIk9crnkYhq41TkwxlZkyn/Fn1YdBgx0Lv5AMjEJpBAgi+FX/pHznr2wM9UqAAO+vCh4rXVGwGCppMvg7uieX6uf8VjTDIMwo8MmuwIDAQAB
-----END PUBLIC KEY-----`

const pubPEM = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrEmbiMh6DoLNg3+xnAwvdQkXRgZBjxgH2XfGhoxXw5w31bDVTUVqiyeHrn4O5tZQW+3X0dIk9crnkYhq41TkwxlZkyn/Fn1YdBgx0Lv5AMjEJpBAgi+FX/pHznr2wM9UqAAO+vCh4rXVGwGCppMvg7uieX6uf8VjTDIMwo8MmuwIDAQAB`

var txt = "{\"tagid\":\"1.305.29.1\",\"userid\":\"2c71488c-8c0c-4fcf-b369-0815841abd5f\",\"model\":\"SM-G965F\",\"country\":\"HK\"}"

func main() {
	
	block, _ := pem.Decode([]byte(PUBLIC_KEY2))
	if block == nil || block.Type != "PUBLIC KEY" {
		fmt.Println(11111111)
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(11111111)
	}

	// Assert that the public key is an *rsa.PublicKey
	rsaPubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		fmt.Println(11111111)
	}

	ciphertext, err := rsaEncryptNoPadding([]byte(txt), rsaPubKey)
	if err != nil {
		panic(err)
	}

	outStr := base64.StdEncoding.EncodeToString(ciphertext)

	fmt.Println(222222222, url.QueryEscape(outStr))

	//plaintext, err := rsaDecryptNoPadding(ciphertext, &rsaPrivateKey)
	//if err != nil {
	//	panic(err)
	//}
	//println(string(plaintext))
}