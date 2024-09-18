package main

import (
	"github.com/DeanThompson/zhihu-go"
	"log"
)

func main() {
	zhihu.Init("./zhihu/config.json")
	topic := zhihu.NewTopic("https://www.zhihu.com/topic/19552832", "")
	showTopic(topic)
}

func showTopic(topic *zhihu.Topic) {
	log.Println("Topic fields:")
	// 链接：https://www.zhihu.com/topic/19552832
	log.Printf("	url: %s", topic.Link)

	// 名称：Python
	log.Printf("	name: %s", topic.GetName())

	// 描述：Python 是一种面向对象的解释型计算机程序设计语言，在设计中注重代码的可读性，同时也是一种功能强大的通用型语言。
	log.Printf("	description: %s", topic.GetDescription())

	// 关注者数量：82805
	log.Printf("	followers num: %d", topic.GetFollowersNum())

	for i, author := range topic.GetTopAuthors() {
		log.Printf("	top-%d author: %s", i+1, author.String())
	}
}
