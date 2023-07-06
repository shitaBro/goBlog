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
	db.Find(&Comment{}).Where("article_id = ?",id).Where("status = ?",1).Count(&totoal)
	err := db.Model(&Comment{}).Limit(size).Offset((page - 1)*size).Order("Created_At DESC").Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?",id).Where("status = ?",1).Scan(&commentList).Error
	if err != nil {
		return commentList,0,errmsg.ERROR
	}
	return commentList,totoal,errmsg.SUCCESS
}

// 删除评论
func DeleteComment(id uint) int {
	var comment Comment
	err := db.Where("id = ?",id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 通过评论
func CheckComment(id int ,data *Comment) int {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status
	err := db.Model(&comment).Where("id = ?",id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id = ?",res.ArticleId).UpdateColumn("comment_count",gorm.Expr("comment_count + ?",1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
 }

 func UnCheckComment(id int,data *Comment) int {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status
	err = db.Model(&comment).Where("id = ?",id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id = ?",res.ArticleId).UpdateColumn("comment_count",gorm.Expr("comment_count - ?",1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
 }