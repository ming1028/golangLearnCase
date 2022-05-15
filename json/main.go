package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	user := Person{
		Name:   "test",
		Weight: 20,
	}
	// 字段没有值时, 默认输出字段的类型零值, 忽略没有值的字段，添加omitempty， json忽略输出: -
	userJson, _ := json.Marshal(user)
	fmt.Println(userJson, string(userJson))

	// 不修改原结构体忽略空值字段
	user1 := Person{
		Name:   "张三",
		Weight: 444,
	}
	user1Json, _ := json.Marshal(PubPerson{Person: &user1})
	fmt.Println(string(user1Json))

	jsonStr1 := `{"id": "1234567","score": "88.50"}` // string id score convert to int cannot unmarshal string into Go struct field Card.id of type int64
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) //

	// json没有整型浮点型 都是number json中的数字经过序列化都是float64
	map1 := make(map[string]interface{}, 1)
	map1["cnt"] = 1
	map1["count"] = "string"
	map1Json, _ := json.Marshal(map1)
	fmt.Printf("str:%#v\n", string(map1Json))
	var map2 map[string]interface{}
	_ = json.Unmarshal(map1Json, &map2)
	fmt.Printf("value:%v\n", map2["cnt"]) // 1
	fmt.Printf("type:%T\n", map2["cnt"])  // float64

	// json decoder
	var map3 map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader(map1Json))
	decoder.UseNumber()
	_ = decoder.Decode(&map3)
	fmt.Println(map3)
	fmt.Printf("value:%v\n", map3["cnt"]) // 1
	fmt.Printf("type:%T\n", map3["cnt"])  // json.Number
	// 将m2["count"]转为json.Number之后调用Int64()方法获得int64类型的值
	count, _ := map3["cnt"].(json.Number).Int64()

	fmt.Printf("type:%T\n", int(count)) // int

	// 匿名结构体添加字段
	json1, _ := json.Marshal(struct {
		*Person
		Token string
	}{
		&user1,
		"lsdjkldsjflsdfj",
	})
	fmt.Println(string(json1))
}

type PubPerson struct {
	*Person
	Name *struct{} `json:"name,omitempty"`
}

type Person struct {
	Name   string   `json:"name"`
	Weight float64  `json:"-"`
	Hobby  []string `json:"hobby,omitempty"`
	Other
	Other2 Other `json:"other2"`
}

type Other struct {
	Other1 string `json:"other1"`
}

type Card struct {
	ID    int64   `json:"id,string"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}
