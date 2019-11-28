package PracticeItem

import (
	"github.com/fpay/foundation-go/job"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Times        int    `yaml:"resendtiems"`
	GrpcPort     string `yaml:"grpc_port"`
	GrpcHost     string `yaml:"grpc_host"`
	RabbitmqUser string `yaml:"mq_user"`
	RabbitmqPwd  string `yaml:"mq_passwd"`
	RabbitmqHost string `yaml:"mq_host"`
	RabbitmqPort int    `yaml:"mq_port"`
	Rabbitmqmail string `yaml:"rabbitmq_mail_queue"`
	Mysqldsn     string `yaml:"mysql_dsn"`
}

func (a *AppConfig) Load() {
	err := viper.Unmarshal(a)
	if err != nil {
		log.Fatalf("failed to parse config file: %s", err)
	}
}
func (a *AppConfig) JobOptions() job.JobManagerOptions {
	return job.JobManagerOptions{
		Username: a.RabbitmqUser,
		Password: a.RabbitmqPwd,
		Host:     a.RabbitmqHost,
		Port:     a.RabbitmqPort,
		Vhost:    a.RabbitmqHost,
	}
}

type AppConfigService interface {
	Load()
	JobOptions()job.JobManagerOptions
}
