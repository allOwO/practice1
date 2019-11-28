package cmd

import (
	"PracticeItem"
	"PracticeItem/pb"
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func init() {
	clientCmd := &cobra.Command{
		Use:   "server",
		Short: "Messenger server",
		Run: func(cmd *cobra.Command, args []string) {
			env := new(PracticeItem.AppConfig)
			env.Load()

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			//建立一个客户端连接,带context
			cc, err := grpc.DialContext(ctx, env.GrpcHost+env.GrpcPort, grpc.WithInsecure())
			log.Println("grpc send host", env.GrpcHost+env.GrpcPort)
			handleInitError(err, "connect")
			defer cc.Close()
			//新建接口
			client := pb.NewMessengerClient(cc)
			//调用gRPC接口
			stream, err := client.SendMessage(context.Background())
			handleInitError(err, "server err")
			defer stream.CloseSend()

			waitc := make(chan struct{})
			// 从客户端得到回信
			go func() {
				for {
					getinfo, err := stream.Recv()
					if err == io.EOF {
						close(waitc)
						return
					}
					if err != nil {
						log.Fatalf("Failed to receive a note : %v", err)
					}
					//返回reponse
					ok := getinfo.GetOk()
					info := getinfo.GetInfo()
					fmt.Println("ok~~!")
					if ok {
						log.Println("ok ,处理成功")
					} else {
						log.Println("通知失败", info)
					}
				}
			}()
			go func() {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					group, messbody, err := getGroupAndMess(scanner.Text())
					if err != nil {
						log.Println(err)
						continue
					}
					log.Println(scanner.Text(), group, messbody)
					err = stream.Send(&pb.Message{
						MessageBody: messbody,
						Group:       group,
					})
					if err != nil {
						log.Fatalf("Failed to send message to sender: %v", err)
					}
				}
			}()
			log.Println("server started")

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			//不明觉厉
		loop:
			for {
				select {
				case <-waitc:
					break loop
				case <-quit:
					break loop
				}
			}
			log.Println("client exited")
		},
	}
	rootCmd.AddCommand(clientCmd)
}

func getGroupAndMess(text string) (string, string, error) {
	text = strings.TrimSpace(text)
	list := strings.Split(text, "#")
	if len(list) != 2 || list[0] == "" || list[1] == "" {
		return "", "", errors.New("Please set true messages")
	}
	return list[0], list[1], nil

}
