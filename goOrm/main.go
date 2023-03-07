package main

import (
	"database/sql"
	"fmt"
	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// UserInfo
type UserInfo22 struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`       // 设置字段大小为255
	MemberNumber *string `gorm:"unique;"`        // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`              // 忽略本字段
}

func main() {
	db, _ := gorm.Open("mysql", "root:root@(localhost)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	db.SingularTable(false) // 不要复数表名
	db.AutoMigrate(&UserInfo22{})

	u1 := UserInfo22{Name: "七米", Age: sql.NullInt64{Int64: 20, Valid: true}, Email: "篮球"}
	u2 := UserInfo22{Name: "七米2", Age: sql.NullInt64{Int64: 20, Valid: true}, Email: "篮球"}

	// 创建记录
	db.Create(&u1)
	db.Create(&u2)

	// 查询
	var u = new(UserInfo22)
	db.First(u)
	fmt.Println("=============")
	fmt.Printf("%#v\n", u)
	user1Map4 := structs.Map(u)
	for k, v := range user1Map4 {
		fmt.Println(k, v, fmt.Sprintf("%T", v))
	}
	fmt.Println("=============")
	var uu UserInfo22
	db.Find(&uu, "email=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("Email", "篮球")
	// 删除
	db.Delete(&u)
}
