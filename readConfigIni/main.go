package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	cfg, err := ini.Load("./readConfigIni/my.ini")
	if err != nil {
		log.Fatalf("read config file err:%+v\n", err)
	}
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("App Mode:", cfg.Section("paths").Key("data").String())

	// 范围限制，以及默认值设置
	fmt.Println("Server Protocol:",
		cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))

	// 类型转换 默认值
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(999))

	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// 修改并保存到文件中
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("./readConfigIni/my.ini.local")
}
