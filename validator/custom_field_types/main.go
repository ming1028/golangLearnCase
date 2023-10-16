package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type DbBackedUser struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt64  `validate:"required"`
}

func main() {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	x := DbBackedUser{
		Name: sql.NullString{String: "1", Valid: true},
		Age:  sql.NullInt64{Int64: 1, Valid: true},
	}
	err := validate.Struct(x)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
		fmt.Println(err)
	}
	return nil
}
