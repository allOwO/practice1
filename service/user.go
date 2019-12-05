package service

import (
	"PracticeItem"
	"github.com/fpay/foundation-go/database"
	"log"
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
func (d *DBservice)GetAllUsers(tablename string) (*[]PracticeItem.User,bool) {
	result := &[]PracticeItem.User{}
	if e := d.Db.Debug().Table("all_users").Where(tablename+" = 1").Find(result).Error; e!=nil {
		return nil,false
	}
	return result,true
}
//更新分组
func (d *DBservice)UpdateUser(tablenames []string, key *PracticeItem.User) bool {
	key = setGroups(tablenames, key)
	if d.Db.Debug().Table("all_users").Where("user_mail = ?",key.Mail).Find(key).RecordNotFound(){
		return false
	}
	if e := d.Db.Debug().Table("all_users").Where("id = ?",key.ID).Updates(key).Error; e!=nil {
		log.Println(e)
		return false
	}
	return true
}
//新建用户
func (d *DBservice)CreateUser(tablenames []string, key *PracticeItem.User) bool {
	key = setGroups(tablenames, key)
	if e := d.Db.Table("all_users").Create(key).Error; e!=nil {
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