package encrypt

import (
	"encoding/base64"
	"strings"
)

//加密
func Encrypt(s, key string) (string, error){
	b, err := AesEncrypt(s, key)
	if err != nil {
		return "", err
	}

	return Base64UrlSafeEncode(b), nil
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

//解密
func Decrypt(data string, key string) (string, error){

	d, _ := Base64URLDecode(data)

	b, err := AesDecrypt(d, key)
	if err != nil {
		return "", err
	}
	return b, nil
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