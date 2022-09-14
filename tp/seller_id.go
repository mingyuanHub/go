package main

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"sort"
	"strconv"
	"fmt"
)

func main() {
	cidList := []int{312,502,726,1267,314,546,870,999}

	for _, cid := range cidList {
		sid := getSupplyChainSidByCid(cid)

		fmt.Println(fmt.Sprintf("%d => %s", cid, sid))
	}
}



func getSupplyChainSidByCid(cid int) string {
	var (
		key1 int
		key2 int
	)

	reg := regexp.MustCompile(`[1-9]+`)
	matches := reg.FindAllString(Md5(strconv.Itoa(cid)), -1)
	matchesInt := StrArrToIntArr(matches)
	sort.Sort(sort.Reverse(sort.IntSlice(matchesInt)))

	key1 = matchesInt[0]
	key2 = cid * 7 + 852

	supplyChainSid := fmt.Sprintf("%d-%d", key1, key2)

	return supplyChainSid
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func StrArrToIntArr(strArr []string) []int {
	var intArr []int
	for _, item := range strArr {
		if i ,err := strconv.Atoi(item); err == nil {
			intArr = append(intArr, i)
		}
	}
	return intArr
}
