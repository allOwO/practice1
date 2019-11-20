package cmd

import (
	"PracticeItem/Globalvar"
	"PracticeItem/controllers"
	"PracticeItem/pb"
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func init() {
	//接收grpc，调用的函数->生产者
	send := &cobra.Command{
		Use: "sender",
		Run: func(cmd *cobra.Command, args []string) {
			port:=viper.GetString("grpc_port")
			if port=="" {
				port = ":8888"
			}
			//固定到8888端口
			listen,err:=net.Listen("tcp",":8888")
			handleInitError(err,"net")
			//新建grpc
			grpcSer:=grpc.NewServer(
				//无应答的存活时间
				grpc.KeepaliveParams(keepalive.ServerParameters{
					Time: 10 * time.Minute,
				}),
			)
			//注册到grpc
			newMess:=controllers.MessageInfo{}
			pb.RegisterMessengerServer(grpcSer,&newMess)
			//注册反射服务，https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch4-rpc/ch4-08-grpcurl.html
			reflection.Register(grpcSer)
			//端口监听交给gprc
			go grpcSer.Serve(listen)
			//退出
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit
			// Finish after all clients disconnected
			log.Println("recv exited")
		},
	}
	recv := &cobra.Command{
		Use: "notification",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
			queue, err := ch.QueueDeclare("send", true, false, false, false, nil)
			if err != nil {
				log.Fatalln(err)
			}
			//给线程设置一个关闭，好像没啥用
			ctx, cencal := context.WithCancel(context.Background())
			for i := 0; i < Globalvar.Recvnum; i++ {
				go controllers.RecvRabbitmq(ctx,queue,ch)
			}
			//退出
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit
			cencal()
			// Finish after all clients disconnected
			log.Println("recv exited")
		},
	}
	jobs:=&cobra.Command{
		Use:"jobs",
	}
	rootCmd.AddCommand(jobs)
	jobs.AddCommand(send)
	jobs.AddCommand(recv)
	jobs.PersistentFlags().IntVarP(&Globalvar.Recvnum, "number", "n", 10, "")
	send.PersistentFlags().StringVarP(&Globalvar.Typ, "type", "t", "mail", "")

}

