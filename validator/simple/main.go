package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	Gender         string     `validate:"oneof=male female prefer_not_to"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	validateStruct()
	validateVariable()
}

func validateStruct() {
	address := &Address{
		Street: "shang hai",
		Planet: "earth",
		Phone:  "1234234234",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Gender:         "male",
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}
	// user := "sfsdfs"
	err := validate.Struct(user)
	if err != nil {
		// 无效参数错误-非struct
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return
	}
}

func validateVariable() {
	myEmail := "ssdfsd.gmail.com"
	errs := validate.Var(myEmail, "required,email")
	if errs != nil {
		fmt.Println(errs)
		return
	}
}
