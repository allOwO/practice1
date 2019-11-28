package controllers

import (
	"PracticeItem"
	"PracticeItem/pb"
	"context"
	jsoniter "github.com/json-iterator/go"
	foundation "github.com/fpay/foundation-go"
	"github.com/fpay/foundation-go/job"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"io"
	"log"
)

type producerController struct {
	user      PracticeItem.UserService
	mess      PracticeItem.MessService
	jobManager foundation.JobManager
	env PracticeItem.AppConfigService
	message   string
	groupName string
	typ       string
	Attempt   int
}

func NewproducerController(user PracticeItem.UserService, mess PracticeItem.MessService,env PracticeItem.AppConfigService, tye string, Attemptnum int) *producerController {
	jobs:=job.NewJobManager(env.JobOptions())
	return &producerController{
		user:    user,
		mess:    mess,
		jobManager:jobs,
		typ:     tye,
		Attempt: Attemptnum,
	}
}
//本来想分开grpc和生产，图方便就这样了

//grpc->生产者
func (p *producerController) SendMessage(stream pb.Messenger_SendMessageServer) error {
	for {
		//接收消息
		log.Println("sendMessage start")
		in, err := stream.Recv()
		log.Println("sendMessage", in, err)
		//流的context,有点不懂？
		if err == io.EOF {
			//链接断开
			return nil
		}
		log.Println("sendMessage no EOF")
		if err != nil {
			log.Println("server lostconnetion")
			return err
		}
		//在之后添加功能
		log.Println(in.GetGroup(), in.GetMessageBody())
		userlist, ok := p.user.GetAllUsers(in.GetGroup())
		if ok == false {
			log.Println("GET Users Fail")
			continue
		}
		messtmp := &PracticeItem.Message{
			Mess:        in.GetMessageBody(),
			Group:       in.GetGroup(),
			Success:     false,
			SuccessUser: 0,
			FailUser:    0,
		}
		p.producer(*userlist, messtmp)
	}
}

type MessJsonBody struct {
	User    PracticeItem.User
	Typ     string
	Message PracticeItem.Message
}

//生产者
func (p *producerController) producer(Users []PracticeItem.User, mbody *PracticeItem.Message) {
	//新建一个队列

	// 发布一个任务

	jobmess := &MessJsonBody{
		User:    item,
		Message: *mbody,
		Typ:     p.typ,
	}
	p.jobManager.Dispatch(context.Background(), jobmess)


	queuename := viper.GetString("rabbitmq_mail_queue")
	dsn := viper.GetString("rabbitmq_dsn")
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatalf("rabbitmq error:%v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Create rabbitmq channal:%V", err)
	}
	queue, err := ch.QueueDeclare(queuename, true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//消息保存到数据库
	p.mess.CreateMessage(mbody)
	for _, item := range Users {
		mess, _ := jsoniter.Marshal(MessJsonBody{
			User:    item,
			Message: *mbody,
			Typ:     p.typ,
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

