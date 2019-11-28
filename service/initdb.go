package service

import (
	"PracticeItem"
)
//初始化表
func (d *DBservice)CheckTables() {
	//有主键索引，用不到其他索引了
	if d.Db.HasTable("all_users") == false {
		d.Db.Table("all_users").CreateTable(&PracticeItem.User{})
	}
	if d.Db.HasTable("messages") == false {
		d.Db.Table("messages").CreateTable(&PracticeItem.Message{})
	}
}
