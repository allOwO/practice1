package model

import (
	"PracticeItem/Globavar"
	"log"
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	slice:=&[]Globavar.User{}
	FindAllUsers("system_user",slice)
	log.Println(slice)
}