package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	UserName string `json:"userName" validate:"required"`
	PassWord string `json:"passWord" validate:"required,min=6,max=20"`
}

func main() {
	example := User{
		PassWord: "123",
	}
	// 中文翻译器
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	// 实例化验证器
	validate := validator.New()
	// 注册翻译器到校验器
	err := zh2.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
		return
	}
	errs := validate.Struct(example)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			fmt.Println(err)
			fmt.Println(err.Translate(trans))
			fmt.Println()
		}
	}
}
