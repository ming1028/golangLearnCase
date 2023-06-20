package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	id   int
	name string
}

func (self *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func main() {
	u := User{1, "tom"}
	fmt.Printf("User: %p, %v\n", &u, u)

	mv := User.TestValue
	mv(u) // 复制reciver

	mp := (*User).TestPointer
	mp(&u) // 指针

	mp2 := (*User).TestValue
	mp2(&u)

	t1 := T{1}
	t2 := &t1
	fmt.Printf("t2 is : %v\n", t2)
	t1.testT()
	t1.testP()
	str := `[{"data": "{\"title\":\"测试视频\",\"description\":\"大数据综合诊断模型，为您的投资保驾护航\",\"url\":\"https://biying-dist.oss-cn-shanghai.aliyuncs.com/dist/third.mp4\",\"picurl\":\"\",\"mediaId\":\"_JwPcekwg9NuWj8fnzZ9f6UkcHS0OavTxje0OEYsAhYLpGWqIr9njqjgZBsDjAPC\",\"thumbMediaId\":\"\"}", "sort": 0, "type": "video"}]`
	ss := make([]map[string]string, 0, 8)
	_ = json.Unmarshal([]byte(str), &ss)
	fmt.Println(ss)
	sss := make(map[string]string)
	_ = json.Unmarshal([]byte(ss[0]["data"]), &sss)
	fmt.Println(sss)
	ssb, _ := json.Marshal(sss)
	fmt.Println(ssb)
}

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("类型 *T 方法集包含全部 receiver T 方法。")
}

func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}
