package service

import (
	"PracticeItem"
	"github.com/fpay/foundation-go/database"
)

type DBservice struct{
	Db *database.DB
}

func NewDBservice(db *database.DB)*DBservice{
	d:=&DBservice{Db:db}
	d.CheckTables()
	return d
}
//查询分组内所有用户
func (d *DBservice)GetAllUsers(tablename string) (result *[]PracticeItem.User,ok bool) {
	result = &[]PracticeItem.User{}
	if b := d.Db.Debug().Table("all_users").Where(tablename+" = 1").Find(result).GetErrors(); len(b)!=0 {
		ok=false
		return
	}
	ok=true
	return
}
//更新分组
func (d *DBservice)UpdateUser(tablenames []string, key *PracticeItem.User) bool {
	key = setGroups(tablenames, key)
	if b:=d.Db.Table("all_users").Where("user_mail = ?",key.Mail).Find(key).GetErrors();len(b)!=0{
		return false
	}
	if b := d.Db.Debug().Table("all_users").Where("id = ?",key.Mail).Updates(key).GetErrors(); len(b)!=0 {
		return false
	}
	return true
}
//新建用户
func (d *DBservice)CreateUser(tablenames []string, key *PracticeItem.User) bool {
	key = setGroups(tablenames, key)
	if b := d.Db.Table("all_users").Create(key).GetErrors(); len(b)!=0 {
		return false
	}
	return true
}

func setGroups(tablenames []string, key *PracticeItem.User) *PracticeItem.User {
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