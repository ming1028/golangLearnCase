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

	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseBool("3"))

	fmt.Println(strconv.FormatInt(-2, 16))
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatUint(7, 4))
	fmt.Println(strconv.FormatFloat(3.1415, 'f', -1, 64))
}
