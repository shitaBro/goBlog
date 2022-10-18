package model

import (
	"fmt"
	"goblog/utils/errmsg"
	
	// "github.com/jinzhu/gorm"
)

type Profile struct {
	// gorm.Model
	ID     int    `gorm:"primarykey" json:"id"`
	User   User   `json:"user"`
	Name   string `gorm:"type:varchar(50)" json:"name"`
	Motto  string `gorm:"type:varchar(255)" json:"motto"`
	Desc   string `gorm:"type:varchar(1024)" json:"desc"`
	Qqchat string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat string `gorm:"type:varchar(200)" json:"wechat"`
	Email  string `gorm:"type:varchar(200)" json:"email" validate:"email"`
	Gitee  string `gorm:"type:varchar(1024)" json:"gitee"`
	Bili   string `gorm:"type:varchar(1024)" json:"bili"`
	Img    string `gorm:"type:varchar(1024)" json:"img"`
	Avatar string `gorm:"type:varchar(1024)" json:"avatar"`
}

func GetProfile(id int) (Profile, int) {
	var profile Profile
	err := db.Where("id = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	db.Select("id").Where("id = ?", id).First(&profile)
	fmt.Println("db .id:", profile.ID)
	
	
	if profile.ID > 0 {
		err := db.Model(&profile).Where("id = ?", id).Update(&data).Error
		if err != nil {
			fmt.Println("update file err:", err)
			return errmsg.ERROR
		}
	} else {
		db.Create(&data)
	}

	return errmsg.SUCCESS
}
