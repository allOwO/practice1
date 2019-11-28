package service

import (
	"PracticeItem"
	"log"
)

func (d *DBservice)GetUserInfo(usermail string) *PracticeItem.WebUserMess {
	senduser := PracticeItem.NewWebUserMess()
	tmp := &PracticeItem.User{}
	selectflag := true
	if ok := d.Db.Table("all_users").Where("user_mail = ?", usermail).First(tmp).RecordNotFound(); ok == false && selectflag {
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