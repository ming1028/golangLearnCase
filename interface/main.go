package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		v    interface{}
		data *byte
		in   interface{}
	)
	v = (*int)(nil)
	fmt.Printf("%#v\n", v)
	fmt.Println(v == nil)

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)

	in = data
	fmt.Println(in, in == nil)
	fmt.Printf("%#v\n", in)
	// 利用反射判断是否为nil
	vi := reflect.ValueOf(in)
	if vi.Kind() == reflect.Ptr {
		fmt.Println("is nil:", vi.IsNil())
	} else {
		fmt.Println("vi not nil")
	}

	/**
	*interface 并不是一个指针类型，包含类型、值
	*包含两类数据结构
	*①runtime.iface 包含方法的的接口
	*②runtime.eface 不包含任何方法的空接口
	 */

	x := interface{}(nil)
	fmt.Printf("%#v\n", x)
	fmt.Println(reflect.TypeOf(x))
	ss, ok := x.(interface{})
	fmt.Println(ok, ss)
}
