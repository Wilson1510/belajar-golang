package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	var config *viper.Viper = viper.New()

	config.SetConfigType("json")
	config.SetConfigName("app")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(config.GetString("name"))
	fmt.Println(config.GetInt("age"))
	fmt.Println(config.GetBool("single"))

	config.SetConfigType("yml")
	config.SetConfigName("config")
	config.AddConfigPath(".")

	err = config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(config.GetString("name"))
	fmt.Println(config.GetInt("age"))
	fmt.Println(config.GetBool("single"))

	config.SetConfigType("env")
	config.SetConfigName("hmm")
	config.AddConfigPath(".")

	err = config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(config.GetString("NAMA"))
	fmt.Println(config.GetInt("AGE"))
	fmt.Println(config.GetString("Email"))

	config.AutomaticEnv()
	fmt.Println(config.GetString("password"))

}