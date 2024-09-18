package main

import (
	"log"
	"os"
)

func main() {
	// 打开 PDF 文件
	filePath := "./colly/files/zy.pdf"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening PDF file: %v\n", err)
	}
	defer f.Close()

	/*// 提取文本
	ctx, err := api.ReadContextFile(filePath)
	if err != nil {
		log.Fatalf("Error reading PDF context: %v\n", err)
	}

	for i := 1; i <= ctx.PageCount; i++ {
		text, err := api
		if err != nil {
			log.Fatalf("Error extracting text from page %d: %v\n", i, err)
		}
		fmt.Printf("Page %d:\n%s\n", i, text)
	}*/
}
