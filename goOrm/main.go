package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
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

	db.LogMode(true)
	// db.Callback().Query().After("gorm:query").Register("my_plugin:after_query", afterQueryCallBack)
	/*db.SingularTable(false) // 不要复数表名
	db.AutoMigrate(&UserInfo22{})

	u1 := UserInfo22{Name: "七米", Age: sql.NullInt64{Int64: 20, Valid: true}, Email: "篮球"}
	u2 := UserInfo22{Name: "七米2", Age: sql.NullInt64{Int64: 20, Valid: true}, Email: "篮球"}
	*/
	/*// 创建记录
	db.Create(&u1)
	db.Create(&u2)*/

	// 查询
	var u = new(UserInfo22)
	/*err := db.Model(u).Where("age = ?", 23).Unscoped().Scan(u).Error
	fmt.Println(err, u.Age)*/
	db.Unscoped().First(u)
	/*fmt.Println("=============")
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
	db.Model(&u).Update("Email", "篮球")*/
	// 删除
	// db.Delete(&u)
	/*us := make([]UserInfo22, 0, 8)
	db.Find(&us).Where("age > ?", 1)*/

	var ages []uint64
	db.Model(&UserInfo22{}).Unscoped().Pluck("age", &ages)
	fmt.Println(ages)

	var us []UserInfo22
	db.Table("user_info22").Unscoped().Where(map[string]interface{}{
		"age": 23,
	}).Find(&us)
	fmt.Println(us)

	db.Where(&UserInfo22{Age: sql.NullInt64{
		Int64: 23,
		Valid: true,
	}}).Unscoped().Find(&us)
	fmt.Println(us)

	db.Not("name = ?", "a").Unscoped().Find(&us)
	fmt.Println(us)

	var u2 = new(UserInfo22)
	db.Set("gorm:query_option", "for update").Unscoped().First(u2)
	fmt.Println(u)

	db.Unscoped().First(u2)
	fmt.Println(u2)

	var u3 UserInfo22
	db.Unscoped().FirstOrInit(&u3, UserInfo22{
		Name: "a",
	})
	fmt.Println(u3)

	var u4 UserInfo22
	db.Unscoped().Where(UserInfo22{
		Name: "c",
	}).Attrs(UserInfo22{Address: "ddd"}).FirstOrCreate(&u4)
	fmt.Println(u4)

	db.Unscoped().Where("name = ?", db.Table("users_goose").Select("MAX(name)").SubQuery()).First(u4)
	fmt.Println(u4)
}

func afterQueryCallBack(scope *gorm.Scope) {
	log.Printf("raw sql:%v vars:%v", scope.SQL, scope.SQLVars)
}
