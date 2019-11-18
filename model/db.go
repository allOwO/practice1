package model

import (
	"PracticeItem/Globavar"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)

var db *gorm.DB

func GetAllKeys(tablename string,slice interface{})error{
	if checkkey(tablename)==true{
		Globavar.MysqlSync.RLock()
		db.Table(tablename).Find(slice)
		Globavar.MysqlSync.RUnlock()
		return nil
	}
	return errors.New("Table Name Error")
}
func InsertKey(tablename string,key interface{})error{
	if checkkey(tablename)==true{
		Globavar.MysqlSync.Lock()
		db.Table(tablename).Save(key)
		Globavar.MysqlSync.Unlock()
		return nil
	}
	return errors.New("Table Name Error")
}
func SetKeys(tablename string,key interface{})error{
	if checkkey(tablename)==true{
		Globavar.MysqlSync.Lock()
		db.Table(tablename).Create(key)
		Globavar.MysqlSync.Unlock()
		return nil
	}
	return errors.New("Table Name Error")
}
//success bug?
func InsertMessSucess(tablename string,message *Globavar.Message)error{
	if checkkey(tablename)==true{
		Globavar.MysqlSync.Lock()
		db.Table(tablename).Where("id = ?", message.ID).UpdateColumn("success_user",gorm.Expr("success_user + ?", 1))
		Globavar.MysqlSync.Unlock()
		return nil
	}
	return errors.New("Table Name Error")
}
func InsertMessFail(tablename string,message *Globavar.Message)error{
	if checkkey(tablename)==true{
		Globavar.MysqlSync.Lock()
		db.Table(tablename).Where("id = ?", message.ID).Update("fail_user",gorm.Expr("fail_user + ?", 1))
		Globavar.MysqlSync.Unlock()
		return nil
	}
	return errors.New("Table Name Error")
}
func checkkey(tablename string)bool{
	if tablename!="system_users" && tablename!="workers"&& tablename!="service_staffs"{
		return false
	}
	return true
}
func init() {
	//config
	user:=viper.GetString("mysql_user")
	passwd:=viper.GetString("mysql_passwd")
	host:=viper.GetInt("mysql_host")
	port:=viper.GetString("mysql_port")
	dbname:=viper.GetString("mysql_dbname")
	//inti mysql
	mysqlArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, passwd, host, port, dbname)
	DB, err := gorm.Open("mysql", mysqlArgs)
	if err != nil {
		log.Fatalf("Open mysql error:%v", err)
	}
	db = DB
	CheckTables()
}


//type User struct {
//	ID        int    `gorm:"AUTO_INCREMENT"`
//	Name      string `gorm:"user_name"`
//	CreatedAt string `gorm:"created_at"`
//	UpdatedAt string `gorm:"updated_at"`
//}
//type Group struct {
//	ID        int    `gorm:"AUTO_INCREMENT"`
//	GroupName string `gorm:"group_name"`
//	CreatedAt string `gorm:"created_at"`
//	UpdatedAt string `gorm:"updated_at"`
//}
//type message struct {
//	ID        int    `gorm:"AUTO_INCREMENT"`
//	Mess      string `gorm:"mess"`
//	GroupID   int    `gormL:"group_id"`
//	CreatedAt string `gorm:"created_at"`
//	UpdatedAt string `gorm:"updated_at"`
//}

func CheckTables() {
	//有主键索引，用不到其他索引了
	if db.HasTable("system_users") == false {
		db.Table("system_users").CreateTable(&Globavar.User{})
	}
	if db.HasTable("workers") == false {
		db.Table("workers").CreateTable(&Globavar.User{})
	}
	if db.HasTable("service_staffs") == false {
		db.Table("service_staffs").CreateTable(&Globavar.User{})
	}
	if db.HasTable("groups") == false {
		db.Table("groups").CreateTable(&Globavar.User{})
	}
	if db.HasTable("messages") == false {
		db.Table("messages").CreateTable(&Globavar.User{})
	}
}
