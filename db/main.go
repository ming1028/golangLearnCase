package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangLearnCase/db/model"
	"time"
)

var db *sql.DB

func main() {
	dsn := "root:root@tcp(127.0.0.1)/sql_test"
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(20)
	db.SetConnMaxLifetime(time.Minute)

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(model.Insert(db))

	fmt.Println(model.Update(db, 444, "王五"))
	fmt.Println(model.Del(db, 1))
}
