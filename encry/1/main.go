package main
import(
	"bytes"
	"crypto/aes"
	"fmt"
	"crypto/cipher"
	"encoding/base64"
	//"net/url"
	"strings"
	"errors"
)
func main() {
	orig := "1000"
	key := "2a98373aacf69156"
	fmt.Println("原文：", orig)
	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：" , encryptCode, len(encryptCode))

	urlen := Base64UrlSafeEncode(encryptCode)
	fmt.Println("Base64安全Url：" , urlen )

	value, _ := Base64URLDecode(urlen)
	fmt.Println("Base64安全unUrl：" , value )
	//
	decryptCode, _ := AesDecrypt(value, key)
	fmt.Println("解密结果：", decryptCode)
}
//加密算法-AES加密
func AesEncrypt(orig string, key string) []byte {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := NewECBEncrypter(block)
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return cryted
}

func AesDecrypt(crytedByte []byte, key string) (string, error) {
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	if len(orig)%blockSize != 0 {
		return "", errors.New(fmt.Sprintf("fail to Notice, CryptBlocks"))
	}

	// 加密模式
	blockMode := NewECBDecrypter(block)

	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	var err error
	orig, err = PKCS7UnPadding(orig)
	if err != nil {
		return "", err
	}
	return string(orig), nil
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length < 1 {
		return []byte{}, errors.New(fmt.Sprintf("fail to Notice, EcpmAesDecrypt1"))
	}
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte{}, errors.New(fmt.Sprintf("fail to Notice, EcpmAesDecrypt2"))
	}
	return origData[:(length - unpadding)], nil
}


type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

//returns a BlockMode which encrypts in electronic code book
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

//returns a BlockMode which decrypts in electronic code book
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}




//base64解码
func Base64URLDecode(data string) ([]byte, error){
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	res, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//安全Base64编码
func Base64UrlSafeEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}