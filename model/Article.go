package model

import (
	"goblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`

	Title   string `gorm:"type:varchar(1024);not null" json:"title"`
	Flag string	`gorm:"type:varchar(20);not null" json:"flag"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(1024)" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
}
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}