package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

type s struct {
	Data map[string]interface{}
}

type stu struct {
	Name string
	Age  int32
}

func main() {
	s1 := s{
		Data: make(map[string]interface{}),
	}
	s1.Data["count"] = 1
	s1Json, _ := json.Marshal(s1.Data)
	fmt.Println(string(s1Json))

	s2 := s{
		Data: make(map[string]interface{}),
	}
	_ = json.Unmarshal(s1Json, &s2.Data)
	fmt.Println(s2)
	for _, v := range s2.Data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
	stu1 := new(stu)
	stu1.Name = "abc"
	stu1.Age = 22
	stu1Json, _ := json.Marshal(*stu1)
	fmt.Println(string(stu1Json))

	map1 := make(map[string]interface{})
	map1["count"] = 1
	map1Json, _ := json.Marshal(map1)
	fmt.Println(string(map1Json))

	s3 := s{
		Data: make(map[string]interface{}),
	}
	s3.Data["cnt"] = 5
	buf := new(bytes.Buffer)
	// 将压缩的数据调用Write方法写入到bytes.Buffer
	enc := gob.NewEncoder(buf)
	_ = enc.Encode(s3.Data)
	s3Json := buf.Bytes()
	fmt.Println(s3Json)

	s4 := s{
		Data: map[string]interface{}{},
	}
	dec := gob.NewDecoder(bytes.NewBuffer(s3Json))
	_ = dec.Decode(&s4.Data)
	fmt.Println(s4)
	for _, v := range s4.Data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}

	stu2 := stu{
		Name: "ss",
		Age:  222,
	}

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	_ = encoder.Encode(stu2)
	stu2Bytes := buffer.Bytes()
	fmt.Println(stu2Bytes)

	stu3 := new(stu)
	decoder := gob.NewDecoder(bytes.NewBuffer(stu2Bytes))
	_ = decoder.Decode(stu3)
	fmt.Println(*stu3)
}
