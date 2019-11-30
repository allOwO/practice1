package controllers

import (
	"PracticeItem"
	"PracticeItem/pb"
	"context"
	"github.com/fpay/foundation-go"
	"github.com/fpay/foundation-go/job"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"io"
	"log"
)

type producerController struct {
	user       PracticeItem.UserService
	mess       PracticeItem.MessService
	jobManager foundation.JobManager
	env        PracticeItem.AppConfigService
	message    string
	groupName  string
	typ        string
	consumernum    int
}

func NewproducerController(user PracticeItem.UserService, mess PracticeItem.MessService, env PracticeItem.AppConfigService, tye string, consumernum int) *producerController {
	//先添加一个队列，防止报错
	dsn:=env.JobOptions()
	addNewQueue(tye,dsn.BuildURL())
	jobs := job.NewJobManager(dsn)
	return &producerController{
		user:       user,
		mess:       mess,
		jobManager: jobs,
		typ:        tye,
		consumernum:    consumernum,
		env:        env,
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
		//
		//if ok == false {
		//		//	log.Println("GET Users Fail")
		//		//	continue
		//		//}
		messtmp := &PracticeItem.Message{
			Mess:        in.GetMessageBody(),
			Group:       in.GetGroup(),
			Success:     false,
		}
		p.producer(in.GetGroup(), messtmp)
	}
}

type MessJsonBody struct {
	Typ       string //邮件，短信，微信公众号
	User      PracticeItem.User  //用户，第二个队列使用
	Message   PracticeItem.Message //消息本体
	delay     int	//重试延迟时间
	quename string	//其余时间和type相同，为了第一次发送消息的队列添加这个字段，详见下方producer
}
//jobManager确定发送队列
func (m *MessJsonBody) Queue() string {
	return m.quename
}
func (m *MessJsonBody) Delay() int {
	return m.delay
}
func (m *MessJsonBody) Marshal() ([]byte, error) {
	return jsoniter.Marshal(m)
}

//生产者
func (p *producerController) producer(Group string, mbody *PracticeItem.Message) {
	// 发布一个任务
	//这里固定了一下队列名称
	log.Println("producer start")
	jobmess := &MessJsonBody{
		Message:   *mbody,
		Typ:       p.typ,
		delay:     10,
		quename: QueueName,
	}
	//检查一下输入错误
	if JobGroupCheck(jobmess)==false{
		log.Println("group error",jobmess.quename,jobmess.Message.Group)
		return
	}
	if JobtypCheck(jobmess)==false{
		log.Println("mess type error",jobmess.Typ,jobmess.Message.Group)
		return
	}
	log.Println("producer send start,queue name",jobmess.Queue())
	err := p.jobManager.Dispatch(context.Background(), jobmess)
	if err != nil {
		log.Println("Producer Send Fail",err)
	}

}
func JobGroupCheck(mess *MessJsonBody) bool {
	if mess.Message.Group == "worker" || mess.Message.Group == "system_user" || mess.Message.Group == "service_staff" {
		return true
	}
	return false
}
func addNewQueue(queuename string,dsn string){
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatalf("rabbitmq error:%v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	_, err = ch.QueueDeclare(queuename, true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
}