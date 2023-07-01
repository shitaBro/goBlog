package model

import (
	"goblog/utils/errmsg"

	"gorm.io/gorm"
)
type Comment struct {
	gorm.Model
	UserId uint `json:"user_id"`
	ArticleId uint `json:"article_id"`
	Title string `json:"article_title"`
	UserName string `json:"username"`
	Content string `gorm:"type:varchar(500):not null;" json:"content"`
	Status int8 `gorm:"type:tinyint:default:2" json:"status"`

}
//添加评论
func AddComment(data *Comment)int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
//根据 [id] 查询单个评论
func GetComment(id int) (Comment,int) {
	var comment Comment
	err := db.Where("id = ?",id).First(&comment).Error
	if err != nil {
		return comment,errmsg.ERROR
	}
	return comment,errmsg.SUCCESS
}
// 获取评论列表 
func GetCommentList(id int,size int,page int,) ([]Comment,int64,int) {
	var commentList []Comment
	var totoal int64
}