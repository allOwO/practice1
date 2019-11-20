package model

import (
	"PracticeItem/Globalvar"
	"log"
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	slice:=&[]Globalvar.User{}
	FindAllUsers("system_user",slice)
	log.Println(slice)
}

func TestUpdateUser(t *testing.T) {
	tmp:=&Globalvar.User{
		Name:         "lll",
		Phone:        "12245678901",
		Mail:         "1@qq.com",
	}
	ok:=UpdateUser([]string{"worker"},tmp)
	log.Println(ok,tmp)
}