package service

import (
	"github.com/fpay/foundation-go/database"
	"log"
	"testing"
)

func TestDBservice_GetAllUsers(t *testing.T) {
	db, err := database.NewDatabase(database.DatabaseOptions{Driver: "mysql",Dsn:"root:12345678@tcp(127.0.0.1:3306)/practice?charset=utf8&parseTime=True&loc=Local"})
	if err!=nil{
		log.Println(err)
	}
	user:=NewDBservice(db)
	log.Println(user.GetAllUsers("worker"))
}