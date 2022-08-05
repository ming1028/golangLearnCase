package main

import (
	"fmt"
)

func main() {
	arr := []int32{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(arr)
	fmt.Println(HeapSort(arr))
}

func HeapSort(arr []int32) []int32 {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, int32(lastlen))
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}

func HeapSortMax(arr []int32, length int32) {
	if length <= 1 {
		return
	}
	depth := length/2 - 1 //二叉树深度
	for i := depth; i >= 0; i-- {
		topmax := i //假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限
			topmax = rightchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
}
