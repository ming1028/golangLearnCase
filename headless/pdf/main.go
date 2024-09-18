package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	httpGet()
	return
	// PDF 文件 URL
	url := "https://www.oajrc.org/FileUpload/PdfFile/29fdc80aeaf440739677bf43a2512c66.pdf"

	// 创建一个上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 创建一个新的 Chrome 实例，并禁用 PDF 查看器插件
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("plugins.plugins_disabled", true),
		chromedp.Flag("disable-pdf-viewer", true),
		chromedp.Flag("no-sandbox", true), // 添加 no-sandbox 标志
		chromedp.Flag("headless", true),
	)
	ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	/*// JavaScript 代码，用于在页面上加载 PDF.js 并解析 PDF 文件
	jsCode := `
	(async function() {
		const pdfjsLib = window['pdfjs-dist/build/pdf'];
		pdfjsLib.GlobalWorkerOptions.workerSrc = 'https://cdnjs.cloudflare.com/ajax/libs/pdf.js/2.9.359/pdf.worker.min.js';

		const loadingTask = pdfjsLib.getDocument('` + url + `');
		const pdf = await loadingTask.promise;
		let content = '';

		for (let i = 1; i <= pdf.numPages; i++) {
			const page = await pdf.getPage(i);
			const textContent = await page.getTextContent();
			const textItems = textContent.items;
			let pageText = '';

			for (let j = 0; j < textItems.length; j++) {
				pageText += textItems[j].str + ' ';
			}

			content += pageText + '\\n';
		}

		return content;
	})()
	`*/

	// 运行任务
	var pdfContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second), // 等待页面加载
	)
	if err != nil {
		log.Fatal(err)
	}

	// 打印 PDF 内容
	fmt.Println("PDF Content:")
	fmt.Println(pdfContent)
}

func httpGet() {
	baseURL := `http://192.168.60.84:3000/html`
	u, _ := url.Parse(baseURL)
	query := u.Query()
	query.Set("url", "https://www.oajrc.org/FileUpload/PdfFile/29fdc80aeaf440739677bf43a2512c66.pdf")
	query.Set("wait_for", "1000")
	query.Set("userAgent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	client := http.Client{}
	request, _ := http.NewRequest("GET", "http://192.168.60.84:3000/html", nil)
	request.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(request)
	fmt.Println(resp)
}
