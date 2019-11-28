package cmd

import (
	"PracticeItem"
	"PracticeItem/controllers"
	"PracticeItem/service"
	"context"
	"github.com/fpay/foundation-go/database"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//启动消息处理，消费通知队列
var notification = &cobra.Command{
		Use: "notification",
		Run: func(cmd *cobra.Command, args []string) {
			env:=new(PracticeItem.AppConfig)
			env.Load()
			db, err := database.NewDatabase(database.DatabaseOptions{Driver: "mysql",Dsn:env.Mysqldsn})
			if err!=nil{
				log.Fatalln("Mysql error:",err)
			}
			user:=service.NewDBservice(db)
			newMess:=controllers.NewMessageInfo(user,user,nil,Typ,Attempt)


			conn, err := amqp.Dial(env.Mysqldsn)
			if err != nil {
				log.Fatalf("rabbitmq error:%v", err)
			}
			defer conn.Close()

			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("Create rabbitmq channal:%V", err)
			}
			//消息公平分派
			ch.Qos(1,0,true)
			//新建一个队列
			queue, err := ch.QueueDeclare(env.Rabbitmqmail, true, false, false, false, nil)
			if err != nil {
				log.Fatalln(err)
			}

			//给线程设置一个关闭，好像没啥用
			ctx, cencal := context.WithCancel(context.Background())
			for i := 0; i < Attempt; i++ {
				go newMess.RecvRabbitmq(ctx,queue,ch)
			}
			//退出
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit
			cencal()
			// Finish after all clients disconnected
			log.Println("notification exited")
		},
	}
func init() {
	jobs.AddCommand(notification)
}