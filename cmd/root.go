package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "messenger",
	Short: "this is a prictice by lzx",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	//设置一个flag
	//cfgFile 设置值，默认值，用法(help)
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config.yaml", "config file path")
}

//初始化配置文件
//
func initConfig() {
	if cfgFile==""{
		//新建yaml文件
		viper.AddConfigPath(".")
		viper.SetConfigName("config.yaml")
		log.Fatalln("No Config File.")
	}
	viper.SetConfigFile(cfgFile) //从string获取文件路径，viper不检查默认
	//获取环境变量
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	//读取
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Read Config File :", viper.ConfigFileUsed())
	} else {
		log.Fatalln("Read Config error",err)
	}
}

func handleInitError(err error, module string) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
