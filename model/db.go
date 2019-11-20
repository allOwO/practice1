package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"PracticeItem/Globalvar"
)

var db *gorm.DB

func FindAllUsers(tablename string,slice *[]Globalvar.User) bool {
	if b := db.Table("all_users").Where(tablename+"= 1").Find(slice).GetErrors(); len(b)!=0 {
		return false
	}
	return true
}

//Update User
//save 更新所有字段
func UpdateUser(tablenames []string, key *Globalvar.User) bool {
	key = setGroups(tablenames, key)
	if b:=db.Table("all_users").Where("user_mail = ?",key.Mail).Find(key).GetErrors();len(b)!=0{
		return false
	}
	if b := db.Debug().Table("all_users").Where("id = ?",key.Mail).Updates(key).GetErrors(); len(b)!=0 {
		return false
	}
	return true
}
//Create user

func CreateUser(tablenames []string, key *Globalvar.User) bool {
	key = setGroups(tablenames, key)
	if b := db.Table("all_users").Create(key).GetErrors(); len(b)!=0 {
		return false
	}
	return true
}
func CreateMessage(key *Globalvar.Message) bool {
	if b := db.Table("messages").Create(key).GetErrors(); len(b)!=0 {
		return false
	}
	return true
}
func setGroups(tablenames []string, key *Globalvar.User) *Globalvar.User {
	su:=false
	w:=false
	ss:=false
	for _, name := range tablenames {
		if name == "system_user" {
			su= true
		} else if name == "worker" {
			w= true
		} else if name == "service_staff" {
			ss = true
		}
	}
	key.SystemUser=su
	key.Workers=w
	key.ServiceStaff=ss
	return key
}
func ChangeGroups(tablenames []string, key interface{}) bool {
	for _, name := range tablenames {
		if b := db.Table(name).Create(key).GetErrors(); b != nil {
			return false
		}
	}
	return true
}

func GetUserInfo(usermail string) *Globalvar.WebUserMess {
	senduser := Globalvar.NewWebUserMess()
	tmp := &Globalvar.User{}
	selectflag := true
	if ok := db.Table("all_users").Where("user_mail = ?", usermail).First(tmp).RecordNotFound(); ok == false && selectflag {
		senduser.Groups = append(senduser.Groups, "system_users")
		selectflag = false
	}
	if tmp.Workers == true {
		senduser.Groups = append(senduser.Groups, "workers")
	}
	if tmp.SystemUser == true {
		senduser.Groups = append(senduser.Groups, "system_user")
	}
	if tmp.ServiceStaff == true {
		senduser.Groups = append(senduser.Groups, "service_staff")
	}
	senduser.UserName = tmp.Name
	senduser.UserMail = tmp.Mail
	senduser.UserPhone = tmp.Phone
	log.Println("find user sql",tmp,senduser)
	return senduser
}

//success bug?
func InsertMessSucess(message *Globalvar.Message) bool {
	if err := db.Table("all_users").Where("id = ?", message.ID).UpdateColumn("success_user", gorm.Expr("success_user + ?", 1)).GetErrors(); err != nil {
		return false
	}
	return true

}
func InsertMessFail(message *Globalvar.Message) bool {
	if err := db.Table("all_users").Where("id = ?", message.ID).UpdateColumn("fail_user", gorm.Expr("fail_user + ?", 1)).GetErrors(); err != nil {
		return false
	}
	return true

}

//func checkkey(tablename string) bool {
//	if tablename != "system_users" && tablename != "workers" && tablename != "service_staffs" {
//		return false
//	}
//	return true
//}
func init() {
	//inti mysql
	mysqlArgs := "root:12345678@(localhost:3306)/practice?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open("mysql", mysqlArgs)
	if err != nil {
		log.Fatalf("Open mysql error:%v", err)
	}
	db = DB
	CheckTables()
}

func CheckTables() {
	//有主键索引，用不到其他索引了
	if db.HasTable("all_users") == false {
		db.Table("all_users").CreateTable(&Globalvar.User{})
	}
	if db.HasTable("messages") == false {
		db.Table("messages").CreateTable(&Globalvar.User{})
	}
}
