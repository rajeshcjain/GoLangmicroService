package main

import (
	"github.com/spf13/viper"
	"log"
	"fmt"
)

func readConfig(){
	viper.SetConfigFile("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()

	if err != nil {
		log.Panic(err)
	}

	/*Reading the file from configuration*/
	server = viper.GetString("development.redis.server")
	port = viper.GetString("development.redis.port")
	fmt.Println("server ",server, " port ",port)
}