package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "100"
	str2 := "100str"
	int1, err := strconv.Atoi(str1)
	fmt.Println(int1, err)
	int2, err := strconv.Atoi(str2)
	fmt.Println(int2, err)
	fmt.Println(strconv.Itoa(100))

	fmt.Println(strconv.ParseBool("TrUE"))
	fmt.Println(strconv.ParseBool("3"))

	fmt.Println(strconv.FormatInt(-2, 2))
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatUint(7, 4))
	fmt.Println(strconv.FormatFloat(3.1415, 'f', -1, 64))

	// base：字符串数字进制 bitSize转换后最大范围
	fmt.Println(strconv.ParseInt("0101", 0, 32))
}
