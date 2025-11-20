package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"bytes"
)

func main() {
	plainText := "118001231"
	eKey := "ca25345f79cb449c9eded9bb459f7304"

	// 加密
	encryptedText, err := encryptAES(plainText, eKey)
	if err != nil {
		fmt.Printf("Encryption error: %v\n", err)
		return
	}

	fmt.Printf("Encrypted text: %s\n", encryptedText)

	// 解密验证
	decryptedText, err := decryptAES(encryptedText, eKey)
	if err != nil {
		fmt.Printf("Decryption error: %v\n", err)
		return
	}

	fmt.Printf("Decrypted text: %s\n", decryptedText)
}

// AES 加密函数
func encryptAES(plainText, key string) (string, error) {
	iv := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	plainData := []byte(plainText)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("AES cipher creation error: %v", err)
	}

	// 添加 PKCS7 填充
	plainData = padPKCS7(plainData, aes.BlockSize)

	if len(plainData)%aes.BlockSize != 0 {
		return "", fmt.Errorf("plain data is not a multiple of the block size")
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	cipherData := make([]byte, len(plainData))
	mode.CryptBlocks(cipherData, plainData)

	encryptedText := base64.StdEncoding.EncodeToString(cipherData)
	return encryptedText, nil
}

// AES 解密函数
func decryptAES(encryptedText, key string) (string, error) {
	iv := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	encryptedData, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("Base64 decode error: %v", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("AES cipher creation error: %v", err)
	}

	if len(encryptedData)%aes.BlockSize != 0 {
		return "", fmt.Errorf("encrypted data is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptedData, encryptedData)

	decryptedData, err := unpadPKCS7(encryptedData)
	if err != nil {
		return "", fmt.Errorf("Unpadding error: %v", err)
	}

	return string(decryptedData), nil
}

// PKCS7 填充
func padPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 移除 PKCS7 填充
func unpadPKCS7(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data")
	}

	padding := int(data[len(data)-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding size")
	}

	for i := len(data) - padding; i < len(data); i++ {
		if data[i] != byte(padding) {
			return nil, fmt.Errorf("invalid padding")
		}
	}

	return data[:len(data)-padding], nil
}