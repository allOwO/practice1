package PracticeItem

import "time"

type User struct {
	ID           int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"Column:user_name;not null"`
	Phone        string `gorm:"Column:user_phone;not null"`
	Mail         string `gorm:"Column:user_mail;unique;not null;index"`
	SystemUser   bool   `gorm:"Column:system_user"`
	Workers      bool   `gorm:"Column:worker"`
	ServiceStaff bool   `gorm:"Column:service_staff"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (u *User)SetGroups(tablenames []string) {
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
	u.SystemUser=su
	u.Workers=w
	u.ServiceStaff=ss
}
type UserService interface {
	GetAllUsers(string) (*[]User, bool)
	UpdateUser([]string,*User) bool
	CreateUser([]string,*User) bool
}
