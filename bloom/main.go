package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"strings"
)

// main
// @Desc: 布隆过滤器
func main() {
	n1 := make([]byte, 10)
	binary.BigEndian.PutUint32(n1, 1)
	fmt.Println(n1)

	// 一百万个元素，误判率为百分之1
	filter := bloom.NewWithEstimates(1000000, 0.01)

	find1 := []byte("hello golang")
	find2 := []byte("hello php")
	find3 := []byte("hello go")

	// 添加元素
	filter.Add(find1)
	filter.Add(find2)
	// filter.Add(find3)

	m, n := bloom.EstimateParameters(1000000, 0.01)
	fmt.Println(bloom.EstimateFalsePositiveRate(m, n, 1000000))

	// 判断是否已经存在
	fmt.Println(filter.Test(find1))
	fmt.Println(filter.Test(find2))
	fmt.Println(filter.Test(find3))
	str := `{"age":12,"length":333}`
	u := new(user)
	_ = json.Unmarshal([]byte(str), u)
	fmt.Println(u.Length, u.Age)

	str1 := "abbbbb"
	str2 := bytes.Trim([]byte(str1), "b")
	str3 := strings.Trim(str1, "b")
	fmt.Println(string(str2), str3)
}

type user struct {
	Age    int32 `json:"age"`
	Length int32 `json:"length"`
}
