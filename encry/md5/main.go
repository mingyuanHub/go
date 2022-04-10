package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
)

func BucketMod(deviceId string) int64 {
	hexstr := genMd5(deviceId)
	bigInt := big.NewInt(0)	// init bigInt
	bigInt.SetString(hexstr, 16)		// assign value of hexstr
	ret := bigInt.Mod(bigInt, big.NewInt(100))
	return ret.Int64()
}

func genMd5(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str) //将str写入到w中
	hexstr := hex.EncodeToString(w.Sum(nil))
	return hexstr
}

func main()  {
	id := BucketMod("UID-eba14ba0-4509-42af-8e43-d00f2a6cdaa0")
	fmt.Println(id)


}


