package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-jarvis/cobrautils"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

type student struct {
	Name string `flag:"name" usage:"student name" persistent:"true"`
	Age  int64  `flag:"age" usage:"student age" shorthand:"a"`
}

var rootCmd = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	stu := student{
		Name: "zhangSan",
		Age:  18,
	}
	cobrautils.BindFlags(rootCmd, &stu)
	_ = rootCmd.Execute()
	fmt.Printf("%+v\n", stu)
	var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary
	stuByte, _ := jsoniter.Marshal(&stu)
	fmt.Println(string(stuByte))
	byteJson := `{"name":"zs", "age":[]}`
	var stuInterface interface{}
	jsoniter.Unmarshal([]byte(byteJson), &stuInterface)
	fmt.Println(stuInterface)
	json.Unmarshal([]byte(byteJson), &stuInterface)
	fmt.Println(stuInterface)
}
