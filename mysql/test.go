package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beego?charset=utf8&loc=Local")

	orm.RegisterModel(new(User))

	//orm.RunSyncdb("default", false, true)
}

func main() {

	orm.Debug = true
	o := orm.NewOrm()

	user := User{Name:"hehe"}
	id, err := o.Insert(&user)
	if err != nil {

	}

	fmt.Println(id, user.Id)

	user.Name = "huhu"
	num, err := o.Update(&user)
	fmt.Println((num))

	user2 := User{Id:user.Id}
	err = o.Read(&user2)
	fmt.Println(user2.Name)

	num, err = o.Delete(&user2)
	fmt.Println(num)

	var maps []orm.Params
	o.Raw("select * from user").Values(&maps)
	if err != nil {
		fmt.Println(num)
	}

	for _, term := range maps {
		fmt.Println(term["id"], ":", term["name"])
	}

}