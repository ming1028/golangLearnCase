package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1)/gorm"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
}
