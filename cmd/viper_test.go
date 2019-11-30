package cmd

import (
	"PracticeItem"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestExecute(t *testing.T) {
	cfg:="../config.yaml"

	viper.SetConfigFile(cfg)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err:=viper.ReadInConfig()
	log.Println(err)
	env:=&PracticeItem.AppConfig{}
	env.Load()
	log.Println(env)

}
