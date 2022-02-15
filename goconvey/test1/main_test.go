package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T)  {
	c := Add(3, 5)
	if c != 8 {
		t.Fatal("add function failed")
	}

	if c != 9 {
		t.Fatal("add function failed")
	}
}

func TestAdd2(t *testing.T) {
	Convey("测试啦", t, func() {
		So(Add(3, 5), ShouldEqual, 7)
	})
}