package cmd

import (
	"PracticeItem"
	"PracticeItem/controllers"
	"PracticeItem/pb"
	"PracticeItem/service"
	"github.com/fpay/foundation-go/database"
	"github.com/spf13/cobra"
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
//发送类型
var Typ string
var Attempt int

var send = &cobra.Command{
	Use: "sender",
	Run: func(cmd *cobra.Command, args []string) {
		env:=new(PracticeItem.AppConfig)
		env.Load()
		db, err := database.NewDatabase(database.DatabaseOptions{Driver: "mysql",Dsn:env.Mysqldsn})
		if err!=nil{
			log.Fatalln("Mysql error:",err)
		}
		user:=service.NewDBservice(db)
		newMess:=controllers.NewproducerController(user,user,env,Typ,Attempt)


		listen, err := net.Listen("tcp", env.GrpcPort)
		log.Println("grpc recv port:", env.GrpcPort)
		handleInitError(err, "net")
		//新建grpc
		grpcSer := grpc.NewServer(
			//无应答的存活时间
			grpc.KeepaliveParams(keepalive.ServerParameters{
				Time: 10 * time.Minute,
			}),
		)
		//注册到grpc
		pb.RegisterMessengerServer(grpcSer, newMess)
		//注册反射服务，https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch4-rpc/ch4-08-grpcurl.html
		reflection.Register(grpcSer)
		//端口监听交给gprc
		go grpcSer.Serve(listen)
		//退出
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		// Finish after all clients disconnected
		log.Println("sender exited")
	},
}
var jobs = &cobra.Command{
	Use: "jobs",
}

func init() {
	//消息发送程序，按照type发送信息
	//接收grpc，调用的函数->生产者
	rootCmd.AddCommand(jobs)
	jobs.AddCommand(send)
	jobs.AddCommand(notification)
	jobs.PersistentFlags().IntVarP(&Attempt, "number", "n", 0, "")
	send.PersistentFlags().StringVarP(&Typ, "type", "t", "mail", "")

}
