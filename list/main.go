package main

import (
	"container/list"
	"fmt"
)

func main() {
	listDemo := list.New()
	listDemo.PushBack(1)
	listDemo.PushBack(2)
	listDemo.PushBack(3)
	listDemo.PushBack(4)

	for e := listDemo.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for listDemo.Len() > 0 {
		fmt.Println(listDemo.Back().Value)
		listDemo.Remove(listDemo.Back())
	}
}
