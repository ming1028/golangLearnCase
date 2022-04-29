package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("./go.mod")
	if err != nil {
		_ = fmt.Errorf("open file failed %s", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println("内容", line)
			}
			fmt.Println("over")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Println(line)
	}
	fmt.Println(cast.ToFloat32("你好"))

	// 读取整个文件
	content, err := ioutil.ReadFile("./file/main.go")
	if err != nil {
		_ = fmt.Errorf("readfile failed :%s", err)
		return
	}
	fmt.Println(string(content))

	// 文件写入
	file, err = os.OpenFile("./writeLog.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 066)
	if err != nil {
		_ = fmt.Errorf("openFile failed :%s", err)
		return
	}
	defer file.Close()
	str := "写入文件\n"
	_, _ = file.Write([]byte(str))
	file.WriteString("写入整个字符串\n")

	file, err = os.OpenFile("./bufio.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		_ = fmt.Errorf("bufio open file error: %s", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("第" + cast.ToString(i) + "行\n") // 数据先写入缓存中
	}
	writer.Flush()

	err = ioutil.WriteFile("./ioutil_write.log", []byte("ioutil write file\n"), 0666)
	if err != nil {
		_ = fmt.Errorf("ioutil write file error: %s", err)
		return
	}

	// copy file
	srcFile := "./writeLog.log"
	dstFile := "newCopyFile.log"
	// 只读方式打开源文件
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Errorf("open "+srcFile+" failed: %v", err)
		return
	}
	defer src.Close()

	// 以写创建的模式打开文件
	dst, err := os.OpenFile(dstFile, os.O_CREATE|os.O_CREATE, 0644)
	if err != nil {
		fmt.Errorf("openFile failed：%v", err)
		return
	}
	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil {
		fmt.Errorf("copy file failed:%v", err)
		return
	}
	fmt.Println("赋值字节数：", n)

	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}
