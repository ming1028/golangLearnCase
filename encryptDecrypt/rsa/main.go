package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	pubKey := getKey("./encryptDecrypt/rsa/pub.pem")
	priKey := getKey("./encryptDecrypt/rsa/pri.pem")

	pubBlock, rst := pem.Decode(pubKey)
	fmt.Println(rst)
	pubInterface, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		panic("parsePublickeyErr")
	}
	pub := pubInterface.(*rsa.PublicKey)
	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte("encry"))
	encryptedStrBase64 := base64.URLEncoding.EncodeToString(encryptedStr)
	fmt.Println(encryptedStrBase64)

	srcEncryPtStr, err := base64.URLEncoding.DecodeString(encryptedStrBase64)
	if err != nil {
		panic("decodeError")
	}

	priBlock, _ := pem.Decode(priKey)
	privateKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		panic("parsePrivatekeyErr")
	}
	decryptStr, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, srcEncryPtStr)
	if err != nil {
		panic("decrypt error")
	}
	srcContent := string(decryptStr)
	fmt.Println(srcContent)
}

func getKey(file string) []byte {
	pubfile, err := os.Open(file)
	if err != nil {
		panic("openFile error")
	}
	defer pubfile.Close()

	pubKey := make([]byte, 2048)
	num, err := pubfile.Read(pubKey)
	if err != nil {
		panic("readError")
	}
	publicKey := pubKey[:num]
	return publicKey
}
