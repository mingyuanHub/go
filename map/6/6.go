package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

type dspAccount struct {
	iosMap map[string]int
}

func (m *dspAccount) Clone() *dspAccount {
	n := *m
	n.iosMap = m.cloneStringIntMap(m.iosMap)
	return &n
}

func (m *dspAccount) cloneStringIntMap(data map[string]int) map[string]int {
	var other = make(map[string]int)
	if len(data) > 0 {
		var jsonStr, _ = json.Marshal(data) // 因data原本就是由json解析而来，这里不应报error
		_ = json.NewDecoder(strings.NewReader(string(jsonStr))).Decode(&other)
	}
	return other
}


func main() {
	newDspAccount := &dspAccount{
		iosMap: map[string]int{},
	}

	//m1 := newDspAccount.Clone()
	//
	//m2 := newDspAccount.Clone()
	//
	//m2.iosMap = map[string]int{"2":1}
	//
	//m3 := newDspAccount.Clone()
	//
	//m3.iosMap = map[string]int{"2":1}
	//
	//m4 := newDspAccount.Clone()
	//
	//fmt.Println(m1.iosMap, m2.iosMap, m4.iosMap)






	list := []*dspAccount{
		newDspAccount,
	}

	list[0] = newDspAccount.Clone()

	list1 := []*dspAccount{
		newDspAccount,
	}

	list1[0] = newDspAccount.Clone()

	list[0].iosMap = map[string]int{"2":1}
	fmt.Println(list1[0])
}
