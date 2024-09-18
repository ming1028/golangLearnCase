package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root@root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDb.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDb.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间

	db.AutoMigrate(&Product{})

	// db.Create(&Product{Code: "500", Price: 33})

	var product Product
	//db.First(&product)
	db.First(&product, "code = ?", "666")
	fmt.Println(product.CreatedAt)
	db.Model(&product).Update("price", 500)
}
