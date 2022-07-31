package main

import (
	"fmt"
	"github.com/spf13/viper"
	"something-init/log"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigName("config-dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	//或者直接解析文件
	//env := "dev"
	//viper.SetConfigFile("./config/config-" + env + ".yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}
}

func main() {
	//初始化日志（zerolog）
	log.SetLogLevel()
}
