package controllers

import (
	"PracticeItem/Globavar"
	"PracticeItem/model"
	"PracticeItem/pb"
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"io"
	"log"
	"time"
)

type MessageInfo struct {
	message   string
	groupName string
}
//grpc->生产者
func (m *MessageInfo) SendMessage(stream pb.Messenger_SendMessageServer) error {
	log.Println("sendMessage success")
	for {
		//接收消息
		in, err := stream.Recv()
		//流的context,有点不懂？
		if err == io.EOF {
			//链接断开
			return nil
		}
		if err != nil {
			log.Println("client lostconnetion")
			return err
		}
		//在之后添加功能
		log.Println(in.GetGroup(), in.GetMessageBody())
		userlist:=[]Globavar.User{}
		if ok:=model.FindAllUsers(in.GetGroup(),&userlist);ok==false{
			log.Fatalln("GET Users Fail")
		}
		messtmp:=&Globavar.Message{
			Mess:        in.GetMessageBody(),
			Group:       in.GetGroup(),
			Success:     false,
			SuccessUser: 0,
			FailUser:    0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		sendToRabbitmq(userlist,messtmp,Globavar.Typ)

	}
}


//生产者
func sendToRabbitmq(Users []Globavar.User,mbody *Globavar.Message,typ string){
	//新建一个队列
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("rabbitmq error:%v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Create rabbitmq channal:%V", err)
	}
	queue, err := ch.QueueDeclare("mail", true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
	model.CreateMessage(mbody)
	for _,item:=range Users {
		mess, _ := jsoniter.Marshal(Globavar.MessJsonBody{
			User:    item,
			Message: *mbody,
			Typ:     typ,
		})
		//生产者推送
		err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType: "",
			Body:        mess,
		})

		if err != nil {
			log.Println(err)

		}
	}

}
//消费者
func RecvRabbitmq(ctx context.Context,queue amqp.Queue,ch *amqp.Channel){
	msg, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("go start")
	//d := <-msg
	//log.Printf("%s", d.Body)
	for {
		select {
		case <-ctx.Done():
			return
		case elem:=<-msg:
			usermess:=Globavar.MessJsonBody{}
			jsoniter.Unmarshal(elem.Body,&usermess)
			if usermess.Typ=="mail"{
				err:=ResendMail(usermess.User.Mail,"hello",usermess.Message.Mess)
				if err!=nil{
					ch.Ack(elem.DeliveryTag,false)
					usermess.Message.UpdatedAt=time.Now()
					model.InsertMessSucess(&usermess.Message)
				}else{
					ch.Ack(elem.DeliveryTag,true)
					usermess.Message.UpdatedAt=time.Now()
					model.InsertMessFail(&usermess.Message)
				}
			}else if usermess.Typ=="Phone"{
				SendShortMess()
				ch.Ack(elem.DeliveryTag,true)
				usermess.Message.UpdatedAt=time.Now()
				model.InsertMessSucess(&usermess.Message)
			}else if usermess.Typ=="wechat"{
				SendOfficialAccount()
				ch.Ack(elem.DeliveryTag,true)
				usermess.Message.UpdatedAt=time.Now()
				model.InsertMessSucess(&usermess.Message)
			}
		}
	}
}