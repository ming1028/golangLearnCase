package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"reflect"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type TypeName struct {
	Name string
	Age  int
}

func main() {
	// 方式1  json序列化 数字类型默认转成float64
	user1 := &UserInfo{
		Name: "张三",
		Age:  22,
	}
	user1Json, _ := json.Marshal(user1)
	fmt.Println(string(user1Json))
	var user1Map map[string]interface{}
	_ = json.Unmarshal(user1Json, &user1Map)
	for key, val := range user1Map {
		fmt.Println(key, val, fmt.Sprintf("%T", val))
	}

	// 方式2 通过反射reflect，判断结构体属性的类型
	user1Map2 := make(map[string]interface{})
	v := reflect.ValueOf(user1)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("not struct"))
	}
	t := v.Type()
	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get("json"); tagValue != "" {
			user1Map2[tagValue] = v.Field(i).Interface()
		}
	}
	fmt.Println(user1Map2)
	for key, val := range user1Map2 {
		fmt.Println(key, val, fmt.Sprintf("%T", val))
	}

	// 第三方structs
	user1Map3 := structs.Map(user1)
	for k, v := range user1Map3 {
		fmt.Println(k, v, fmt.Sprintf("%T", v))
	}

	userJson := `{"name":"张三","age":22}`
	user4 := new(UserInfo)
	_ = json.Unmarshal([]byte(userJson), user4)
	fmt.Printf("%v\n", user4)
	u4Val := reflect.ValueOf(user4)
	if u4Val.Kind() == reflect.Ptr {
		u4Val = u4Val.Elem()
	}
	fmt.Println(u4Val)
	u4Type := u4Val.Type()
	for i := 0; i < u4Val.NumField(); i++ {
		field := u4Type.Field(i)
		fmt.Printf("%s, %T\n", field.Tag.Get("json"), u4Val.Field(i).Interface())
	}

	type demoTest struct {
		Add string
		TypeName
	}
	dt := demoTest{
		"234234",
		TypeName{
			Name: "adf",
			Age:  12,
		},
	}
	fmt.Println("================")
	user1Map4 := structs.Map(&dt)
	for k, v := range user1Map4 {
		fmt.Println(k, v, fmt.Sprintf("%T", v))
	}

	// 返回所有value的切片 []interface{}
	fmt.Println(structs.Values(dt))
	fmt.Println(structs.Names(dt))
	fmt.Println("fields", structs.Fields(user1))
	fmt.Println("name", structs.Name(dt))
}
