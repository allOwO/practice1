package PracticeItem

import (
	"github.com/fpay/foundation-go/job"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Times         int    `mapstructure:"attempt_times";yaml:"attempt_times"`
	GrpcPort      string `mapstructure:"grpc_port";yaml:"grpc_port"`
	GrpcHost      string `mapstructure:"grpc_host";yaml:"grpc_host"`
	RabbitmqUser  string `mapstructure:"mq_user";yaml:"mq_user"`
	RabbitmqPwd   string `mapstructure:"mq_passwd";yaml:"mq_passwd"`
	RabbitmqVHost string `mapstructure:"mq_vhost";yaml:"mq_vhost"`
	RabbitmqHost  string `mapstructure:"mq_host";yaml:"mq_host"`
	RabbitmqPort  int    `mapstructure:"mq_port";yaml:"mq_port"`
	Rabbitmqmail  string `mapstructure:"rabbitmq_mail_queue";yaml:"rabbitmq_mail_queue"`
	Mysqldsn      string `mapstructure:"mysql_dsn";yaml:"mysql_dsn"`
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
		Vhost:    a.RabbitmqVHost,
	}
}
func (a *AppConfig) AttemptTimes() int {
	return a.Times
}

type AppConfigService interface {
	Load()
	AttemptTimes() int
	JobOptions() job.JobManagerOptions
}
