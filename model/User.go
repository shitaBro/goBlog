package model

import (
	"encoding/base64"
	"goblog/utils/errmsg"
	"log"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" label:"角色码"`
}

// 登录验证
func CheckLogin(username string, password string) (int, int) {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USERNAME_NOT_EXIST, int(user.ID)
	}
	if ScryptPsw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG, int(user.ID)
	}
	//非管理员 没权限
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT, int(user.ID)
	}
	return errmsg.SUCCESS, int(user.ID)
}

// 查询用户是否存在
func CheckUserExits(username string) int {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //用户存在
	}
	return errmsg.SUCCESS
}

// 更新user前查询
func CheckUpUser(id int, username string) int {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		//用户名已被使用
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPsw(data.Password)
	err := db.Create(data).Error
	if err == nil {
		return errmsg.SUCCESS
	}
	return errmsg.ERROR
}

// 查询单个用户
func GetUserInfo(id int) (User, int) {
	var user User
	err := db.Where("id = ?", id).First(&user).Omit("password").Error
	if err == nil {
		return user, errmsg.SUCCESS
	}
	return user, errmsg.ERROR_USERNAME_NOT_EXIST
}

// 获取用户列表
func GetUserList(username string, page int, size int) ([]User, int64) {
	var users []User
	var totoal int64
	if username != "" {
		db.Select("id,username,role").Where("username LIKE ?", "%"+username+"%").Limit(size).Offset(page * size).Find(&users).Count(&totoal)
		return users, totoal
	}
	db.Select("id,username,role").Limit(size).Offset(page * size).Find(&users).Count(&totoal)
	return users, totoal
}

// 编辑用户
func EditUserInfo(id int, data *User) int {
	var user User
	var mapU = make(map[string]interface{})
	mapU["username"] = data.Username
	mapU["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(&mapU).Error
	if err == nil {
		return errmsg.SUCCESS
	}
	return errmsg.ERROR

}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 重设密码
func ChangePsw(id int, data *User) int {
	err := db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 重置密码
func ResetPsw(id int) int {
	err := db.Model(&User{}).Where("id = ?", id).Update("password", "123456").Error
	if err == nil {
		return errmsg.SUCCESS
	}
	return errmsg.ERROR
}

// 密码加密
func ScryptPsw(password string) string {
	const keyLen = 8
	salt := []byte{5, 48, 12, 33, 85, 6, 10, 9}
	Hashpsw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	finalPsw := base64.StdEncoding.EncodeToString(Hashpsw)
	return finalPsw
}

func (u *User)BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPsw(u.Password)
	u.Role = 2
	return nil
}
func (u *User)BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPsw(u.Password)
	return nil
}
// 前端登录
func CheckLoginFront(username string,psw string) (User,int) {
	var user User
	var pswError error
	db.Where("username = ?",username).Find(&user)
	pswError = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(psw))
	if (user.ID == 0) {
		return user,errmsg.ERROR_USERNAME_NOT_EXIST
	}
	if (pswError != nil) {
		return user,errmsg.ERROR_PASSWORD_WRONG
	}
	return user,errmsg.SUCCESS
}
