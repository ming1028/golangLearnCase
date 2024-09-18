package main

import (
	"fmt"
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
	"regexp"
)

var (
	seg    gse.Segmenter
	posSeg pos.Segmenter

	new, _ = gse.New("zh,testdata/test_en_dict3.txt", "alpha")

	text = "你好世界, Hello world, Helloworld."
)

func main() {
	seg.LoadDict()
	segCut()
}

func cut() {
	hmm := new.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = new.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)
	fmt.Println("analyze: ", new.Analyze(hmm, text))

	hmm = new.CutAll(text)
	fmt.Println("cut all: ", hmm)

	reg := regexp.MustCompile(`(\d+年|\d+月|\d+日|[\p{Latin}]+|[\p{Hangul}]+|\d+\.\d+|[a-zA-Z0-9]+)`)
	text1 := `헬로월드 헬로 서울, 2021年09月10日, 3.14`
	hmm = seg.CutDAG(text1, reg)
	fmt.Println("Cut with hmm and regexp: ", hmm, hmm[0], hmm[6])
}

func segCut() {
	// 分词文本
	tb := []byte("I've got something.")

	// 处理分词结果
	fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.String(string(tb), true))
	fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(string(tb)))

	segments := seg.Segment(tb)
	// 处理分词结果, 普通模式
	fmt.Println(gse.ToString(segments))

	/*segments1 := seg.Segment([]byte(text))
	// 搜索模式
	fmt.Println(gse.ToString(segments1, true))*/
}
