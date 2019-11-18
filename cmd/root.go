package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "messenger",
	Short: "this is a prictice by lzx",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}


func init() {
	//设置一个flag
	//cfgFile 设置值，默认值，用法(help)
	rootCmd.PersistentFlags().StringVarP(&cfgFile,"config","c","/root/Practice/config.yaml","config file path")
	cobra.OnInitialize(initConfig)
	fmt.Println(cfgFile)
}

//初始化配置文件
//
func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) //从string获取文件路径，viper不检查默认
	} else {
		//新建yaml文件，存放在home
		viper.AddConfigPath("./")
		viper.SetConfigName("config.yaml")
		//初始化一些默认值
		//viper.SetDefault("mysqlport","3306")
	}
	//获取环境变量
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	//读取
	if err := viper.ReadInConfig();err == nil {
		log.Println("Read Config File :", viper.ConfigFileUsed())
	} else {
		log.Fatalln("Fail to Read Config File.",err,viper.ConfigFileUsed())
	}
}

func handleInitError(err error, module string) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}