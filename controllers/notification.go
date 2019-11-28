package controllers

import (
	"PracticeItem"
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type consumerController struct {
	mess       PracticeItem.MessService
	mail       PracticeItem.SendMailService
	Createsync sync.Once
	message    string
	groupName  string
}
func NewconsumerController(mess PracticeItem.MessService, mail PracticeItem.SendMailService) *consumerController {
	return &consumerController{
		mess:    mess,
		mail:    mail,
	}
}

//消费者
func (c *consumerController) Consumer(ctx context.Context,) {
	msg, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("recv rabbitmq start,queue name:", queue.Name)
	for {
		select {
		//收到ctrl+c 关闭
		case <-ctx.Done():
			return
		case elem := <-msg:
			//解消息json
			usermess := MessJsonBody{}
			jsoniter.Unmarshal(elem.Body, &usermess)
			if usermess.Typ == "mail" {
				err := ResendMail(usermess.User.Mail, "hello", usermess.Message.Mess)
				if err == nil {
					//true 为之前所有的都确认
					ch.Ack(elem.DeliveryTag, false)
					m.mess.InsertMessSucess(&usermess.Message)
				} else {
					ch.Ack(elem.DeliveryTag, false)
					m.mess.InsertMessFail(&usermess.Message)
				}
			} else if usermess.Typ == "phone" {
				SendShortMess()
				ch.Ack(elem.DeliveryTag, false)
				m.mess.InsertMessSucess(&usermess.Message)
			} else if usermess.Typ == "wechat" {
				SendOfficialAccount()
				ch.Ack(elem.DeliveryTag, false)
				m.mess.InsertMessSucess(&usermess.Message)
			}
		}
	}
}
