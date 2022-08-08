package main

import (
	"fmt"
	"math"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

type Tree struct {
	root   *Node
	length int
}

func main() {
	creatTree()
}

func creatTree() {
	arrList := []int{14, 2, 5, 7, 23, 35, 12, 17, 31}
	myTree := Tree{}
	for i := 0; i < len(arrList); i++ {
		myTree = insertNode(myTree, arrList[i])
		myTree.length++
	}
	fmt.Println(myTree)
	//LDR(myTree)
	TreeHeight(myTree)
}
func TreeHeight(tree Tree) {
	var hl = 1
	if tree.root.left != nil {
		hl = heightMax(tree.root.left, hl)
	}
	var hr = 1
	if tree.root.right != nil {
		hr = heightMax(tree.root.left, hr)
	}
	fmt.Println(hl, hr)
	fmt.Println("Tree height is ", int(math.Max(float64(hl), float64(hr))))
}

func heightMax(node *Node, h int) int {
	var hL = h
	var hR = h
	if node.left == nil && node.right == nil {
		fmt.Println(node)
		return h
	}
	if node.left != nil {
		h++
		hL = heightMax(node.left, h)
	}
	if node.right != nil {
		h++
		hR = heightMax(node.right, h)
	}
	return int(math.Max(float64(hL), float64(hR)))
}

//中序遍历
func LDR(tree Tree) {
	readList := make(map[int]int)
	i := 0
	var currentNode *Node
	currentNode = tree.root
	for {
		//fmt.Println(currentNode)
		if i == tree.length {
			//fmt.Println(currentNode.value)
			break
		}
		if currentNode.left == nil {
			if readList[currentNode.value] == 1 {
				if readList[currentNode.right.value] == 1 {
					currentNode = currentNode.parent
					continue
				} else {
					currentNode = currentNode.right
					continue
				}
			} else {
				fmt.Println(currentNode.value)
				readList[currentNode.value] = 1
				i++
				if currentNode.right == nil {
					currentNode = currentNode.parent
					continue
				} else {
					if readList[currentNode.right.value] == 1 {
						currentNode = currentNode.parent
						continue
					} else {
						currentNode = currentNode.right
						continue
					}
				}
			}
		} else {
			if readList[currentNode.left.value] == 1 {
				if readList[currentNode.value] == 1 {
					currentNode = currentNode.right
					continue
				} else {
					fmt.Println(currentNode.value)
					readList[currentNode.value] = 1
					i++
					if currentNode.right == nil {
						currentNode = currentNode.parent
						continue
					} else {
						if readList[currentNode.right.value] == 1 {
							currentNode = currentNode.parent
							continue
						} else {
							currentNode = currentNode.right
							continue
						}

					}
				}
			} else {
				currentNode = currentNode.left
				continue
			}

		}

	}
}

func insertNode(tree Tree, insertValue int) Tree {
	var currentNode *Node
	var tmp *Node
	i := 0
	if tree.length == 0 {
		currentNode = new(Node)
		currentNode.value = insertValue
		tree.root = currentNode
		return tree
	} else {
		currentNode = tree.root
	}
	for {
		//fmt.Println(currentNode)
		if currentNode.value < insertValue {
			//判断是否有右节点
			if currentNode.right == nil {
				tmp = new(Node)
				tmp.value = insertValue
				currentNode.right = tmp
				tmp.parent = currentNode
				break
			} else {
				currentNode = currentNode.right
				continue
			}
		} else {
			if currentNode.left == nil {
				tmp = new(Node)
				tmp.value = insertValue
				currentNode.left = tmp
				tmp.parent = currentNode
				break
			} else {
				currentNode = currentNode.left
				continue
			}
		}
		i++
	}
	return tree
}
