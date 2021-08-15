package main

import (
	"encoding/xml"
	"fmt"
)
type Student struct{
	Name string
	Address Addr
}
type Addr struct {
	City string `xml:"city"`
	Build string  `xml:"build,attr"`
	MyItem Item
}

type Item struct {
	Attributes []xml.Attr `xml:",any,attr"`
}

func main() {

	p := &Student{Name:"gaofeng", Address:Addr{City:"sz",Build:"2b"}}

	p.Address.MyItem = Item{}
	p.Address.MyItem.Attributes = []xml.Attr{}
	tt:=xml.Attr{Name: xml.Name{ Space:".", Local :"临时1" }, Value:"vvv&11"}
	p.Address.MyItem.Attributes = append(p.Address.MyItem.Attributes, tt)

	tt=xml.Attr{Name: xml.Name{ Space:"", Local :"临时2" }, Value:"vvv22"}
	p.Address.MyItem.Attributes = append(p.Address.MyItem.Attributes, tt)

	// 生成xml
	buf, _ := xml.Marshal(p)
	fmt.Println(string(buf))

	// 解析xml
	var s = `<Student><Name>gaofeng</Name><Address build="2b"><city>sz</city><MyItem 临时1="vvv&amp;11" 临时2="vvv22"></MyItem></Address></Student>`
	fmt.Println(s)
	pp := new(Student)
	xml.Unmarshal([]byte(s), pp)
	fmt.Println(pp.Address.City)
	fmt.Println(pp.Address.MyItem.Attributes[0].Value)
}