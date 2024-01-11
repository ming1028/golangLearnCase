package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		name    string
		age     int32
		married bool
	)
	// 从标准输入扫描文本，读取由空白符分隔的值
	fmt.Scan(&name, &age, &married) // 强制类型转换
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容")
	text, err := reader.ReadString('\n') // 读到换行
	if err != nil {
		fmt.Errorf("read from stdin err:%v", err)
		return
	}
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)

	fmt.Println(fmt.Sprintf("%03s", strconv.Itoa(22)))
}
