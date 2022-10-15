package utils
import (
	"fmt"
	"os"
	"github.com/go-ini/ini"
)
var (
	AppMode  string
	HttpAddress string
	HttpPort string
	JwtKey string

	Db     string
	DbHost string
	DbPort string
	DbUser string
	DbPwd  string
	DbName string

	AccessKey   string
	SecretKey   string
	Bucket       string
	QiniuServer string
)

func init() {
	file,err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路劲",err)
		os.Exit(1)
	}
	LoadServer(file)
	LoadDataBase(file)
	LoadQiniu(file)

}
func LoadServer(file *ini.File) {
	section := file.Section("server")
	AppMode = section.Key("AppMode").MustString("debug")
	HttpAddress = section.Key("HttpAddress").MustString("0.0.0.0")
	HttpPort = section.Key("HttpPort").MustString(":3000")
	JwtKey = section.Key("JwtKey").MustString("4ai7ng2y0use")

}
func LoadDataBase(file *ini.File) {
	section := file.Section("database")
	Db = section.Key("Db").MustString("mysql")
	DbHost = section.Key("DbHost").MustString("localhost")
	DbPort = section.Key("DbPort").MustString(":3000")
	DbUser = section.Key("DbUser").MustString("root")
	DbPwd = section.Key("DbPwd").MustString("12345678")
	DbName = section.Key("DbName").MustString("ginblog")
}
func LoadQiniu(file *ini.File) {
	section := file.Section("qiniu")
	AccessKey = section.Key("AccessKey").String()
	SecretKey = section.Key("SecretKey").String()
	Bucket = section.Key("Bucket").String()
	QiniuServer = section.Key("QiniuServer").String()
}
