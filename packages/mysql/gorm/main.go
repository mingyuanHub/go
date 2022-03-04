package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Db)

	db, err = gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err)
	}

	//设置最大空闲连接
	db.DB().SetMaxIdleConns(10)
	//设置最大连接数
	db.DB().SetMaxOpenConns(100)
	//设置连接超时时间:1分钟
	db.DB().SetConnMaxLifetime(time.Minute)

	//关闭连接
	defer db.Close()

	//sql日志
	db.LogMode(true)

	//db.SetLogger(gorm.Logger{})
	//db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	//ping
	err = db.DB().Ping()
	fmt.Println(err)

	//查询记录
	var (
		articles []*Article
		count int
	)
	db.Find(&articles)
	db.Where("id = 2").Find(&articles)
	db.Debug().Where("id > ?", 2).Or("id = ?",5).Order("id desc").Limit(10).Offset(0).Find(&articles).Count(&count)


	sql := "select * from article"
	db.Exec(sql).Find(&articles)

	//db.Debug().Where(&Article{Id: 5}).Find(&articles)
	fmt.Println(articles, articles[0],count)

	//新增记录
	//err = db.Debug().Table("article").Create(&Article{Id: 5}).Error
	//fmt.Println(err)
	//db.Debug().Model(&Article{}).Create(&Article{Id: 6})

	//更新记录
	//article := &Article{Title:"ttt1"}
	//err = db.Debug().Model(article).Update(article).Error
	//fmt.Println(err)

	//删除记录
	var dCount int
	db.Delete(&Article{Id:2})
	fmt.Println(dCount)

}

func main() {

}
