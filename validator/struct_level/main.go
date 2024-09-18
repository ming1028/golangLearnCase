package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type validationError struct {
	Namespace       string `json:"namespace"`
	Field           string `json:"field"`
	StructNamespace string `json:"structNamespace"`
	StructField     string `json:"structField"`
	Tag             string `json:"tag"`
	ActualTag       string `json:"actualTag"`
	Kind            string `json:"kind"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	Param           string `json:"param"`
	Message         string `json:"message"`
}

type Gender uint

const (
	Male Gender = iota + 1
	Female
	Intersex
)

func (gender Gender) String() string {
	terms := []string{"Male", "Female", "Intersex"}
	if gender < Male || gender > Intersex {
		return "unknown"
	}
	return terms[gender]
}

type User struct {
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email" json:"email"`
	FavouriteColor string     `validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `validate:"required,dive,required"`
	Gender         Gender     `json:"gender" validate:"required,gender_custom_validation"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "_" {
			return ""
		}
		return name
	})

	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	err := validate.RegisterValidation("gender_custom_validation", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(Gender)
		return value.String() != "unknown"
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// build 'User' info, normally posted data etc...
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "Unknown",
	}

	user := &User{
		FirstName:      "",
		LastName:       "",
		Age:            45,
		Email:          "Badger.Smith@gmail",
		FavouriteColor: "#000",
		Addresses:      []*Address{address},
	}

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err = validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			e := validationError{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            fmt.Sprintf("%v", err.Kind()),
				Type:            fmt.Sprintf("%v", err.Type()),
				Value:           fmt.Sprintf("%v", err.Value()),
				Param:           err.Param(),
				Message:         err.Error(),
			}

			indent, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			fmt.Println(string(indent))
		}
		return
	}
}

func UserStructLevelValidation(sl validator.StructLevel) {

	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
		sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
	}
}
