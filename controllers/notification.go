package controllers

import (
	"PracticeItem"
	"context"
	"encoding/json"
	"github.com/fpay/foundation-go"
	"github.com/fpay/foundation-go/job"
	"log"
)

//发送消息 的队列 的名称
const QueueName = "message"

type consumerController struct {
	user       PracticeItem.UserService
	mess       PracticeItem.MessService
	mail       PracticeItem.SendMailService
	jobManager foundation.JobManager
	env        PracticeItem.AppConfigService
	QueueName  string
}

func NewconsumerController(user PracticeItem.UserService, mess PracticeItem.MessService, mail PracticeItem.SendMailService, env PracticeItem.AppConfigService, quename string) *consumerController {
	jobs := job.NewJobManager(env.JobOptions())
	return &consumerController{
		user:       user,
		mess:       mess,
		mail:       mail,
		env:        env,
		jobManager: jobs,
		QueueName:  quename,
	}
}

//消费者
func (c *consumerController) Consumer(ctx context.Context) {
	log.Println("Consumer start,queue name:", QueueName)
	err := c.jobManager.Do(ctx, QueueName, c.jobForMess)
	if err != nil {
		log.Println("Consumer启动失败.")
	}
	//for {
	//	select {
	//	//收到ctrl+c 关闭
	//	case <-ctx.Done():
	//		return
	//	case elem := <-msg:
	//	}
	//}
}

//对队列的每个用户进行处理
func (c *consumerController) UserConsumer(ctx context.Context) {
	log.Println("UserConsumer start,queue name:", c.QueueName)
	err := c.jobManager.Do(ctx, c.QueueName, c.jobForUser)
	if err != nil {
		log.Println("UserConsumer启动失败.")
	}
}

//处理单个消息，把用户发送到队列
func (c *consumerController) jobForMess(ctx context.Context, jobber foundation.Jobber) error {
	jobmess := new(MessJsonBody)
	log.Printf("jobForMess start,%s", string(jobber.Body()))
	err := json.Unmarshal(jobber.Body(), jobmess)
	if err != nil || jobmess.Message.Group == "" {
		return jobber.Skip()
	}
	//把用户放到消息队列
	userlist, ok := c.user.GetAllUsers(jobmess.Message.Group)
	log.Println("userlist", userlist)
	if ok == false {
		log.Println("GET Users Fail")
		return jobber.Skip()
	}
	for _, user := range *userlist {
		tmp := &MessJsonBody{
			Typ:     jobmess.Typ,
			User:    user,
			Message: jobmess.Message,
			delay:   10,
			quename: jobmess.Typ, //队列名和类型相等，在放入队列前检查过了
		}
		log.Println("JobForMess send to ",jobmess.Typ)
		err = c.jobManager.Dispatch(context.Background(), tmp)
		if err != nil {
			log.Println("Send mq Fail")
			if jobber.Attempt() < c.env.AttemptTimes() {
				return jobber.Retry(context.Background(), jobmess.Delay())
			}
		}
	}

	//消息存入mysql
	jobmess.Message.Success = true
	ok = c.mess.CreateMessage(&jobmess.Message)
	if ok == false {
		if jobber.Attempt() < c.env.AttemptTimes() {
			return jobber.Retry(context.Background(), jobmess.Delay())
		}
	}
	return jobber.OK()
}

//接收每个用户的消息，处理
func (c *consumerController) jobForUser(ctx context.Context, jobber foundation.Jobber) error {
	jobmess := new(MessJsonBody)
	err := json.Unmarshal(jobber.Body(), jobmess)
	if err != nil {
		return jobber.Skip()
	}
	log.Println("jobForUser get from ", jobmess.Queue())
	err = c.mail.SendMail(jobmess.User.Mail, "hello", jobmess.Message.Mess)
	if err != nil {
		if jobber.Attempt() < c.env.AttemptTimes() {
			return jobber.Retry(context.Background(), jobmess.Delay())
		} else {
			log.Printf("send to %v Fail:%v", jobmess.User.Mail, err)
		}
	}
	// 处理成功
	return jobber.OK()
}

func JobtypCheck(mess *MessJsonBody) bool {
	if mess.Typ == "mail" || mess.Typ == "short" || mess.Typ == "wechat" {
		return true
	}
	return false
}
