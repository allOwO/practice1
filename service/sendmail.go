package service

import (
	"context"
	"github.com/aymerick/raymond"
	"github.com/mailgun/mailgun-go/v3"
	"io/ioutil"
	"log"
	"os"
	"time"
)
//重试
type sendMailService struct {
	template *raymond.Template
	mail *mailgun.MailgunImpl
	email string
	total string
	messagebody string
}

//读入HTML模板和初始化邮件
func NewsendMailService()*sendMailService{
	tmp:=&sendMailService{}
	bytes,err:=ioutil.ReadFile("index.html")
	if err!=nil ||len(bytes)==0{
		log.Printf("Get Email Template Error:%v",err)
	}
	templateHTML:=string(bytes)
	temp,err:=raymond.Parse(templateHTML)
	if err!=nil{
		log.Println("template error:",err)
		os.Exit(1)
	}
	tmp.template=temp
	//邮件初始化
	tmp.mail = mailgun.NewMailgun(`sandbox2a7b23ca21f642d58385f80df1fad0db.mailgun.org`, `e850f4b368aea62eca0839f2b931905a-1df6ec32-9e34e1ba`)
	return tmp

}
//发邮件
func (s *sendMailService)SendMail(email string, total string, messagebody string) error{
	mess := s.mail.NewMessage(
		"lzx 575361715@qq.com",
		s.total,
		s.messagebody,
		s.email,
	)
	mess.SetHtml(s.getHtmlEmail())
	ctx, can := context.WithTimeout(context.Background(), time.Second*10)
	defer can()
	_, _, err := s.mail.Send(ctx, mess)
	return err
}

func (s *sendMailService)getHtmlEmail()string{
	tmpm:=map[string]string{
		"total":s.total,
		"message":s.messagebody,
	}
	result,err:=s.template.Exec(tmpm)
	if err!=nil{
		log.Println("template error:",err)
		os.Exit(1)
	}
	return result
}
