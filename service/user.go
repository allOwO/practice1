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
	tmp:=&PracticeItem.User{}
	if d.Db.Table("all_users").Where("user_mail = ?",key.Mail).Find(tmp).RecordNotFound(){
		return false
	}
	tmp.SetGroups(tablenames)
	tmp.Phone=key.Phone
	tmp.Name=key.Name
	log.Println("UpdateUser",*tmp)
	if e := d.Db.Debug().Table("all_users").Save(tmp).Error; e!=nil {
		return false
	}
	return true
}
//新建用户
func (d *DBservice)CreateUser(tablenames []string, key *PracticeItem.User) bool {
	key.SetGroups(tablenames)
	if e := d.Db.Table("all_users").Create(key).Error; e!=nil {
		return false
	}
	return true
}

