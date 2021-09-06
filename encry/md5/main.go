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
	id := BucketMod("UID-f3bd5289-7cb2-4a33-8763-d811d3425906")
	fmt.Println(id)
}


