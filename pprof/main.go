package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func test() {

	log.Println(" ===> loop begin.")
	for i := 0; i < 1000; i++ {
		log.Println(genSomeBytes())
	}

	log.Println(" ===> loop end.")
}

//生成一个随机字符串
func genSomeBytes() *bytes.Buffer {

	var buff bytes.Buffer

	for i := 1; i < 20000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}

	return &buff
}

func main() {

	go func() {
		for {
			test()
			time.Sleep(time.Second * 1)
		}
	}()

	//启动pprof
	http.ListenAndServe("0.0.0.0:10000", nil)
}

// 访问localhost:10000/debug/pprof 点击profile生成一个profile文件  使用go tool pprof profile文件路径
// 输入web 生成调用文件图片

// Mac： brew install graphviz

// Windows: 下载graphviz 将graphviz安装目录下的bin文件夹添加到Path环境变量中。 在终端输入dot -version查看是否安装成功
