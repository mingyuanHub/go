package test1

import (
	"sync"
	"testing"
)

var test = map[int]int{1:1}

func BenchmarkNoLock(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		_, _ = test[1]
		_, _ = test[2]
	}
}

var lock sync.Mutex

func BenchmarkLock(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		lock.Lock()
		_, _ = test[1]
		_, _ = test[2]
		lock.Unlock()
	}
}


