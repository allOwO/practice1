package controllers

import (
	"context"
	"github.com/aymerick/raymond"
	"github.com/mailgun/mailgun-go/v3"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestSendMail(t *testing.T) {
	bytes,err:=ioutil.ReadFile("index.html")
	if err!=nil{
		log.Printf("Get Email Template Error:%v",err)
		os.Exit(1)
	}
	templateHTML:=string(bytes)
	template,err=raymond.Parse(templateHTML)
	if err!=nil{
		log.Println("template error:",err)
		os.Exit(1)
	}
	mail:=mailgun.NewMailgun("sandbox2a7b23ca21f642d58385f80df1fad0db.mailgun.org","e850f4b368aea62eca0839f2b931905a-1df6ec32-9e34e1ba")
	mess:=mail.NewMessage(
		"lzx 575361715@qq.com",
		"Hello",
		"this is a test",
		"alrightowo@gmail.com",
	)
	mess.SetHtml(GetHtmlEmail("test","ttt"))

	_,id,err:=mail.Send(context.Background(),mess)
	log.Println("mail id",id)
	if err!=nil{
		log.Println("mail error:",err)
	}
}