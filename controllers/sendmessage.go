package controllers

import (
	"context"
	"errors"
	"github.com/mailgun/mailgun-go/v3"
	"github.com/aymerick/raymond"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"time"
)
var template *raymond.Template
var mail *mailgun.MailgunImpl
var ResendTimes int
//重试
func ResendMail(email string, total string, messagebody string)error{
	times:=viper.GetInt("resendtiems")
	times=limitnum(times)//防止配置次数设置太多
	for ;times>0;times--{
		err:=SendMail(email,total,messagebody)
		if err==nil{
			return nil
		}
	}
	return errors.New("Send Email Fail!")
}
//发邮件
func SendMail(email string, total string, messagebody string) error{
	mess := mail.NewMessage(
		"lzx 575361715@qq.com",
		total,
		messagebody,
		email,
	)
	mess.SetHtml("")
	ctx, can := context.WithTimeout(context.Background(), time.Second*10)
	defer can()
	_, _, err := mail.Send(ctx, mess)
	return err
}

func SendShortMess(){}
func SendOfficialAccount(){}
func GetHtmlEmail(total string,message string)string{
	tmpm:=map[string]string{
		"total":total,
		"message":message,
	}
	result,err:=template.Exec(tmpm)
	if err!=nil{
		log.Println("template error:",err)
		os.Exit(1)
	}
	return result
}
//读入HTML模板和初始化邮件
func init(){
	bytes,err:=ioutil.ReadFile("index.html")
	if err!=nil ||len(bytes)==0{
		log.Printf("Get Email Template Error:%v",err)
		os.Exit(1)
	}
	templateHTML:=string(bytes)
	temp,err:=raymond.Parse(templateHTML)
	if err!=nil{
		log.Println("template error:",err)
		os.Exit(1)
	}
	template=temp
	//邮件初始化
	mail = mailgun.NewMailgun(`sandbox2a7b23ca21f642d58385f80df1fad0db.mailgun.org`, `e850f4b368aea62eca0839f2b931905a-1df6ec32-9e34e1ba`)

}

func limitnum(num int)int{
	if num<0{
		return 1
	}
	if num>=20{
		return 20
	}
	num++//这里是总次数=重试+1
	return num
}
