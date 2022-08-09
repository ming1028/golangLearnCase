package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	userPwd := "123456"
	passwordbyte, err := generatePwd(userPwd)
	if err != nil {
		fmt.Println("加密出错了")
	}
	fmt.Println(string(passwordbyte))

	pwd := "$2y$10$lvXuhX1DyVoGAXBQ/aytHuFUz4hBFPEH58y3E5gtocpjX4O77POJS"
	isOk, _ := validatePwd(userPwd, pwd)
	if !isOk {
		fmt.Println("密码错误")
		return
	}
	fmt.Println(isOk)
}

func generatePwd(userPwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPwd), bcrypt.DefaultCost)
}

func validatePwd(userPwd string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPwd)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}
