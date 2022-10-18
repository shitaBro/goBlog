package model

import (
	"fmt"
	"goblog/utils"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB
var dberr error
func init() {

	// dsn := "root:12345678@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local&interpolateParams=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",utils.DbUser,utils.DbPwd,utils.DbHost,utils.DbPort,utils.DbName)
	db,dberr = gorm.Open(utils.Db,dsn)
	if dberr != nil {
		fmt.Println("数据库连接失败,请检查参数",dberr)
	}else {
		fmt.Println("数据库链接成功")
	}
	
	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)
	//自动迁移
	db.AutoMigrate(&User{},&Profile{})
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

}