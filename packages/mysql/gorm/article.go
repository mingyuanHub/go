package main

type Article struct {
	Id         int    `gorm:"column:id;primaryKey"`
	Title      string `gorm:"column:title"`
	Content    string `gorm:"column:content"`
	Desc       string `gorm:"column:desc"`
	Status     int    `gorm:"column:status"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

func (Article) Table() string {
	return "article"
}
