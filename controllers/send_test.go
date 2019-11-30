package controllers

import (
	"PracticeItem"
	"context"
	"github.com/fpay/foundation-go/job"
	"github.com/fpay/rabbiter"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestS(t *testing.T) {
	cfg := "../config.yaml"
	viper.SetConfigFile(cfg)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	env := new(PracticeItem.AppConfig)
	env.Load()
	log.Println(env.JobOptions())

	jobManager := job.NewJobManager(env.JobOptions())
	jobmess := &MessJsonBody{
		Message: PracticeItem.Message{
			Mess:    "body",
			Group:   "group",
			Success: false,
		},
		Typ:     "mail",
		delay:   10,
		quename: QueueName,
	}
	log.Println("producer send start,queue name", jobmess.Queue())
	err = jobManager.Dispatch(context.Background(), jobmess)
	log.Println("producer send", err)
}

func TestRmq(t *testing.T) {
	r := rabbiter.NewRabbiter(rabbiter.Options{
		Host: "amqp://guest:guest@localhost:5672/",
	})

	p := r.Publisher()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := strconv.FormatInt(int64(time.Now().Nanosecond()), 10)

	err := p.Publish(
		ctx,
		&rabbiter.Exchange{Name: "rabbiter_routing", Type: "direct", AutoDeleted: false, Durable: true},
		&rabbiter.Message{ID: id, Body: []byte("hello"), RoutingKey: "jerray"},
	)
	if err != nil {
		panic(err)
	}
}
func TestOld(t *testing.T) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("rabbitmq error:%v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Create rabbitmq channal:%v", err)
	}
	queue, err := ch.QueueDeclare("mail", true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	mess, _ := jsoniter.Marshal(MessJsonBody{
		User: PracticeItem.User{
			Phone:        "",
			Mail:         "575361715@qq.com",
		},
		Message: PracticeItem.Message{},
		Typ:     "mail",
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
