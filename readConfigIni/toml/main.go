package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./readConfigIni/toml/config.toml")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.GetString("database.server"))
	var conf conf
	_ = viper.Unmarshal(&conf)
	fmt.Println(conf)
}

type db struct {
	Server string
	Ports  []int32
}

type serv struct {
	Ip string
}

type conf struct {
	Serv serv `mapstructure:"servers"`
	Db   db   `mapstructure:"database"`
}
