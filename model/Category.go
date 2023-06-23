package model

import (
	"goblog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func CheckCategoryExits(name string) int {
	var cate Category
	db.Where("name = ?", name).First(&cate)
	if cate.Name != "" {
		return errmsg.ERROR_CATEGORYNAME_USED
	}

	return errmsg.SUCCESS
}
func AddCategory(data *Category) (int, string) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR, err.Error()
	}
	return errmsg.SUCCESS, ""
}
func GetCategory(id int) (Category, int) {
	var cate Category
	err := db.Where("id = ?", id).First(&cate).Error
	if err != nil {
		return cate, errmsg.ERROR
	}
	return cate, errmsg.SUCCESS
}

// 查询分类列表
func GetCategories(name string, pageSize int, pageNum int) ([]Category, int64) {
	var cates []Category
	var totoal int64
	if name != "" {
		db.Where("name = ?", "%"+name+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Count(&totoal)
		return cates, totoal
	}
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	db.Model(&cates).Count(&totoal)
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil, 0
	}
	return cates, totoal
}
func EditCategory(id int, name string) int {
	var cate Category
	cate.Name = name
	err := db.Model(&cate).Where("id = ?", id).Updates(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func DeleteCategory(id int) int {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
