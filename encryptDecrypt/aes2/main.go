package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var (
	AesKey    = []byte("1SFo4Fp7TQ5J757t90Nq62sAD844e247")
	encryData = "qltp3jkH5bcLGdZJoZSAK4/zha8IensAQVdJux+htzPRsIToKiJPGjJsLbA/yO5ZnmMAOCeLxvLog8FKpw3mp0D/chy2crPY9f15KdO9/U0xIyOKeI7wgeEV1jfqRJsRMqHIm7iYug8RcSO3ibYvcg=="
	iv        = "7d92I8zRk3UcO5P5"
	data      = []byte(`{"market":"SZ","accountId":22321,"dealNum":200,"dealPrice":19.19,"stockCode":"002733","tradeType":1}`)
)

func main() {
	block, err := aes.NewCipher(AesKey)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	fmt.Println(blockSize)

	// base64解密
	srcData, _ := base64.StdEncoding.DecodeString(encryData)

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(srcData))
	blockMode.CryptBlocks(origData, srcData)
	padding := int(origData[len(origData)-1])
	orignData := origData[:(len(origData) - padding)] // 根据填充值，获得填充的长度，截取原数据
	fmt.Println(orignData)

	// 加密
	enBlock, _ := aes.NewCipher(AesKey)
	enBlockSize := enBlock.BlockSize()

	// 填充
	lenEnPadding := enBlockSize - (len(data) % enBlockSize) // 填充的长度
	byteEnPadding := byte(lenEnPadding)                     // 填充的字符
	enDataByte := bytes.Repeat([]byte{byteEnPadding}, lenEnPadding)
	encryData := append(data, enDataByte...)
	fmt.Println(encryData)
	blocMode := cipher.NewCBCEncrypter(enBlock, []byte(iv))
	crypted := make([]byte, len(encryData))
	blocMode.CryptBlocks(crypted, encryData)
	base64Str := base64.StdEncoding.EncodeToString(crypted)
	fmt.Println(base64Str)
}
