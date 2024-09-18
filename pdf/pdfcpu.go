package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/unidoc/unipdf/v3/common"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	str := strings.FieldsFunc(`图表目录　　图表光芯片行业现状　　图表光芯片行业产业链调研　　……　　图表2019-2024年光芯片行业市场容量统计　　图表2019-2024年中国光芯片行业市场规模情况　　图表光芯片行业动态　　图表2019-2024年中国光芯片行业销售收入统计　　图表2019-2024年中国光芯片行业盈利统计　　图表2019-2024年中国光芯片行业利润总额　　图表2019-2024年中国光芯片行业企业数量统计　　图表2019-2024年中国光芯片行业竞争力分析　　……　　图表2019-2024年中国光芯片行业盈利能力分析　　图表2019-2024年中国光芯片行业运营能力分析　　图表2019-2024年中国光芯片行业偿债能力分析　　图表2019-2024年中国光芯片行业发展能力分析　　图表2019-2024年中国光芯片行业经营效益分析　　图表光芯片行业竞争对手分析　　图表**地区光芯片市场规模　　图表**地区光芯片行业市场需求　　图表**地区光芯片市场调研　　图表**地区光芯片行业市场需求分析　　图表**地区光芯片市场规模　　图表**地区光芯片行业市场需求　　图表**地区光芯片市场调研　　图表**地区光芯片行业市场需求分析　　……　　图表光芯片重点企业（一）基本信息　　图表光芯片重点企业（一）经营情况分析　　图表光芯片重点企业（一）盈利能力情况　　图表光芯片重点企业（一）偿债能力情况　　图表光芯片重点企业（一）运营能力情况　　图表光芯片重点企业（一）成长能力情况　　图表光芯片重点企业（二）基本信息　　图表光芯片重点企业（二）经营情况分析-9-光芯片行业现状及前景2024-2030年中国光芯片行业研究分析与前景趋势预测报告中国产业调研网　www.cir.cn　　图表光芯片重点企业（二）盈利能力情况　　图表光芯片重点企业（二）偿债能力情况　　图表光芯片重点企业（二）运营能力情况　　图表光芯片重点企业（二）成长能力情况　　……　　图表2024-2030年中国光芯片行业信息化　　图表2024-2030年中国光芯片行业市场容量预测　　图表2024-2030年中国光芯片行业市场规模预测　　图表2024-2030年中国光芯片行业风险分析　　图表2024-2030年中国光芯片市场前景分析　　图表2024-2030年中国光芯片行业发展趋势　　略……订阅“2024-2030年中国光芯片行业研究分析与前景趋势预测报告”，编号：3285050，请致电：4006128668、010-66181099、010-66182099、010-66183099Email邮箱：kf@Cir.cn详细内容：https://www.cir.cn/0/05/GuangXinPianHangYeXianZhuangJiQianJing.html了解更多，请访问上述链接，以下无内容！！-10-光芯片行业分析报告2024-2030年中国光芯片行业研究分析与前景趋势预测报告中国产业调研网　www.cir.cn`, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSpace(r)
	})
	fmt.Println(strings.Join(str, ""))
	// 设置调试模式
	common.SetLogger(common.NewConsoleLogger(common.LogLevelDebug))

	// PDF 文件路径
	filePath := "./colly/files/zy.pdf"

	// 打开 PDF 文件
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("无法打开 PDF 文件: %v", err)
	}
	defer f.Close()

	// 加载 PDF 文档
	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		log.Fatalf("无法读取 PDF 文件: %v", err)
	}

	// 获取 PDF 文档页数
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		log.Fatalf("无法获取 PDF 页数: %v", err)
	}

	page, err := pdfReader.GetPage(numPages)
	if err != nil {
		log.Fatalf("无法获取第 %d 页: %v", numPages, err)
	}

	// 创建文本提取器
	ex, err := extractor.New(page)
	if err != nil {
		log.Fatalf("无法创建提取器: %v", err)
	}

	// 提取文本
	text, err := ex.ExtractText()
	if err != nil {
		log.Fatalf("无法提取第 %d 页的文本: %v", numPages, err)
	}
	fmt.Println(text)

	// 遍历每一页，提取文本内容
	for i := 1; i <= numPages; i++ {
		// 获取页面
		page, err := pdfReader.GetPage(i)
		if err != nil {
			log.Fatalf("无法获取第 %d 页: %v", i, err)
		}

		// 创建文本提取器
		ex, err := extractor.New(page)
		if err != nil {
			log.Fatalf("无法创建提取器: %v", err)
		}

		// 提取文本
		text, err := ex.ExtractText()
		if err != nil {
			log.Fatalf("无法提取第 %d 页的文本: %v", i, err)
		}

		// 打印每一页的文本内容
		fmt.Printf("第 %d 页的内容:\n%s\n", i, text)
	}
}
