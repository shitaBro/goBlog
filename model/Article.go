package model

import (
	"goblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`

	Title   string `gorm:"type:varchar(1024);not null" json:"title"`
	Flag    string `gorm:"type:varchar(20);not null" json:"flag"`
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
func GetOneArticle(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOTEXIST
	}
	return art, errmsg.SUCCESS
}
func GetArticles(keywords string, pageSize int, pageNum int) (articles []Article, code int, num int64) {
	var articleList []Article
	var totoal int64
	if keywords == "" {
		err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Preload("Category").Find(&articleList).Error
		db.Model(&articleList).Count(&totoal)
		if err != nil {
			return nil, errmsg.ERROR, totoal
		}
		return articleList, errmsg.SUCCESS, totoal
	}
	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Preload("Category").Where("title LIKE ?", keywords+"%").Find(&articleList).Error
	db.Model(&articleList).Where("title LIKE ?", keywords+"%").Count(&totoal)
	if err != nil {
		return nil, errmsg.ERROR, totoal
	}
	return articleList, errmsg.SUCCESS, totoal

}
