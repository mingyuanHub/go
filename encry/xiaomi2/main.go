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

const PUBLIC_KEY = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrEmbiMh6DoLNg3+xnAwvdQkXR
gZBjxgH2XfGhoxXw5w31bDVTUVqiyeHrn4O5tZQW+3X0dIk9crnkYhq41TkwxlZk
yn/Fn1YdBgx0Lv5AMjEJpBAgi+FX/pHznr2wM9UqAAO+vCh4rXVGwGCppMvg7uie
X6uf8VjTDIMwo8MmuwIDAQAB
-----END PUBLIC KEY-----`

const pubPEM = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrEmbiMh6DoLNg3+xnAwvdQkXRgZBjxgH2XfGhoxXw5w31bDVTUVqiyeHrn4O5tZQW+3X0dIk9crnkYhq41TkwxlZkyn/Fn1YdBgx0Lv5AMjEJpBAgi+FX/pHznr2wM9UqAAO+vCh4rXVGwGCppMvg7uieX6uf8VjTDIMwo8MmuwIDAQAB`


func encrypt(text string) (string, error) {
	// Decode the public key from Base64  
	//decodedKey, err := base64.StdEncoding.DecodeString(pubPEM)
	//if err != nil {
	//	return "", err
	//}

	// Parse the public key  
	block, _ := pem.Decode([]byte(PUBLIC_KEY))
	if block == nil || block.Type != "PUBLIC KEY" {
		return "", fmt.Errorf("failed to decode PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Assert that the public key is an *rsa.PublicKey  
	rsaPubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("failed to assert public key as *rsa.PublicKey")
	}

	// Encrypt the text using the RSA public key with no padding  
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, []byte(text))
	if err != nil {
		return "", err
	}

	// Encode the encrypted bytes to Base64  
	outStr := base64.StdEncoding.EncodeToString(encrypted)
	return outStr, nil
}

func main() {
	text := "{\"tagid\":\"1.305.29.1\",\"userid\":\"2c71488c-8c0c-4fcf-b369-0815841abd5f\",\"model\":\"SM-G965F\",\"country\":\"HK\"}"
	encrypted, err := encrypt(text)
	if err != nil {
		fmt.Printf("Error encrypting text: %s\n", err)
		return
	}
	fmt.Printf("Encrypted text: %s\n", encrypted)

	fmt.Println(222222222, url.QueryEscape(encrypted))
}