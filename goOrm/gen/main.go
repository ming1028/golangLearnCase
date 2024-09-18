package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	gen2 "gorm.io/gen"
	"gorm.io/gorm"
)

const MySQLDSN = "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True"

func main() {
	gen := gen2.NewGenerator(gen2.Config{
		OutPath: "./goOrm/gen/model",
		Mode:    gen2.WithDefaultQuery | gen2.WithQueryInterface,
	})
	gen.UseDB(connectDB(MySQLDSN))
	gen.ApplyBasic(gen.GenerateAllTable()...)

	gen.Execute()
}

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}
