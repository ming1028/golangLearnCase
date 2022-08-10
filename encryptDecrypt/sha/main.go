package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	str := "sssssjlksdflsdf"
	data := []byte(str)
	has := sha1.Sum(data)
	shaStr1 := fmt.Sprintf("%x", has)
	fmt.Println(shaStr1)

	w := sha1.New()
	io.WriteString(w, str)
	// w.Write(data)
	bw := w.Sum(nil)
	shaStr2 := hex.EncodeToString(bw)
	fmt.Println(shaStr2)

}
