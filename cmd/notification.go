package cmd

import (
	"PracticeItem"
	"PracticeItem/controllers"
	"PracticeItem/service"
	"context"
	"github.com/fpay/foundation-go/database"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)
var Cosumernum int
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
			//给线程设置一个关闭，好像没啥用
			ctx, cencal := context.WithCancel(context.Background())

			user:=service.NewDBservice(db)
			mail:=service.NewsendMailService()
			newMess:=controllers.NewconsumerController(user,user,mail,env,Typ)
			//这个是对单个消息的消费
			go newMess.Consumer(ctx)
			//对每个用户消费
			if Cosumernum<1{Cosumernum=1}
			for i:=0;i<Cosumernum;i++ {
				go newMess.UserConsumer(ctx)
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