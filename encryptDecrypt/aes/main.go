package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

var PwdKey = []byte("DIS**#KKKDJJSKDI")

func main() {
	str := []byte("12321321321")
	pwd, _ := EnPwdCode(str)
	fmt.Println(pwd)
	bytes, _ := DePwdCode(pwd)
	fmt.Println(string(bytes))
}

// PKCS7Padding 填充 需要填充的长度数字n重复 n填充到原数据末尾
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize // 填充的长度
	bytePadding := byte(padding)
	fmt.Println(bytePadding, string(bytePadding)) // uint8对应ascii码表
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误")
	}
	fmt.Println("UnPadding", string(origData), origData[length:])
	unpadding := int(origData[length-1]) //末尾填充内容，填充规则：(末尾填充3 个 3) 获取填充内容就可以得到填充长度
	return origData[:(length - unpadding)], nil
}

func AesEcrypt(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	origData = PKCS7Padding(origData, blockSize)
	fmt.Println(string(origData))
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func EnPwdCode(pwd []byte) (string, error) {
	result, err := AesEcrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

// 解密
func DePwdCode(pwd string) ([]byte, error) {
	//解密base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	//执行AES解密
	return AesDeCrypt(pwdByte, PwdKey)

}

// 实现解密
func AesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 使用密钥作为偏移量
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	fmt.Println(string(origData))
	//去除填充字符串
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}
