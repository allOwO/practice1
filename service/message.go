package service

import (
	"PracticeItem"
	"github.com/jinzhu/gorm"
)

func (d *DBservice)CreateMessage(key *PracticeItem.Message) bool {
	if b := d.Db.Table("messages").Create(key).GetErrors(); len(b)!=0 {
		return false
	}
	return true
}


func (d *DBservice)InsertMessSucess(message *PracticeItem.Message) bool {
	if err := d.Db.Table("messages").Where("id = ?", message.ID).UpdateColumn("success_user", gorm.Expr("success_user + ?", 1)).GetErrors(); err != nil {
		return false
	}
	return true

}
func (d *DBservice)InsertMessFail(message *PracticeItem.Message) bool {
	if err := d.Db.Table("messages").Where("id = ?", message.ID).UpdateColumn("fail_user", gorm.Expr("fail_user + ?", 1)).GetErrors(); err != nil {
		return false
	}
	return true
}