package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 创建文件
	file, _ := os.Create("./fileOperate/create.log")
	defer file.Close()

	// 获取文件信息
	fileInfo, _ := file.Stat()
	fmt.Printf("file name is %s\nPermissions is %s\nModTime is %s\n",
		fileInfo.Name(), fileInfo.Mode(), fileInfo.ModTime())

	// 改变文件权限
	_ = file.Chmod(0777)

	// 改变拥有者
	_ = file.Chown(os.Getuid(), os.Getgid())
	fileInfo, _ = file.Stat()
	fmt.Printf("Permissions is %s\n", fileInfo.Mode())

	// 删除文件
	_ = os.Remove("./fileOperate/create.log")
}

func writeAll(filename string) error {
	err := os.WriteFile(filename, []byte("hello world"), 0666)
	if err != nil {
		return err
	}
	return nil
}

// 直接操作io
func writeLine(filename string) error {
	data := []string{
		"hello",
		"world",
		"golang",
	}
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()

	for _, line := range data {
		_, _ = f.WriteString(line + "\n")
	}
	return nil
}

// 使用缓存区写入
func writeLine2(filename string) error {
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	bufferWriter := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		n, _ := bufferWriter.WriteString("hello world" + cast.ToString(i) + "\n")
		fmt.Println(n)
	}
	// 将内存buffer到硬盘
	_ = bufferWriter.Flush()
	return nil
}

// 偏移量写入
func writeAt(filename string) error {
	data := []byte{
		0x41, // A
		0x73, // s
		0x20, // space
		0x20, // space
		0x67, // g
	}
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	_, _ = file.Write(data)

	replaceSplace := []byte{
		0x6F, // o
		0x6E, // n
	}
	_, _ = file.WriteAt(replaceSplace, 2) // 文件打开不能使用os.O_APPEND
	return nil
}

// 缓存区写入
func writeBuffer(filename string) error {
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	bufferWrite := bufio.NewWriter(file)
	n, _ := bufferWrite.WriteString("1111")
	fmt.Println(n)

	unflushBufferSize := bufferWrite.Buffered()
	fmt.Println("缓存区字节数：", unflushBufferSize)

	bytesAvailable := bufferWrite.Available()
	fmt.Println("缓存区可用字节数：", bytesAvailable)

	_ = bufferWrite.Flush()
	return nil
}

func readAll(filename string) error {
	data, _ := os.ReadFile(filename)
	fmt.Printf("content is %s\n", data)
	return nil
}

func readAll2(filename string) error {
	file, _ := os.Open(filename)
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	log.Printf("content is %s\n", content)
	return nil
}

func readLine(filename string) error {
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	bufferedReader := bufio.NewReader(file)

	for {
		lineBytes, err := bufferedReader.ReadBytes('\n')
		if err == io.EOF {
			return nil
		}
		log.Printf("every line data is %s\n", lineBytes)
	}
}
