package main

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
	"sync"
)

type Item generic.Type

// Node 节点
type Node struct {
	key   int32
	value Item
	left  *Node
	right *Node
}

// ItemBinarySearchTree 树结构
type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func main() {

}

// Insert 节点插入
func (bst *ItemBinarySearchTree) Insert(key int32, value Item) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := &Node{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// insertNode 寻找节点应该插入的位置
func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// InOrderTraverse 中序遍历
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(item Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

// inorderTraverse 按顺序遍历
func inOrderTraverse(n *Node, f func(item Item)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}

// PreOrderTraverse 前序遍历
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

// preOrderTraverse
func preOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		f(n.value)
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}

// PostOrderTraverse 后序遍历
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value)
	}
}

// Min 最小值
func (bst *ItemBinarySearchTree) Min() *Item {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

// Max 最大值
func (bst *ItemBinarySearchTree) Max() *Item {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

// Search 查找节点是否存在
func (bst *ItemBinarySearchTree) Search(key int32) bool {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return search(bst.root, key)
}

func search(n *Node, key int32) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		search(n.left, key)
	}
	if key > n.key {
		search(n.right, key)
	}
	return true
}

// Remove 移除节点
func (bst *ItemBinarySearchTree) Remove(key int32) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

func remove(node *Node, key int32) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}

	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}

	// 1、删除节点右子树的最小值 2、删除节点左子树的最大值
	leftmostrightside := node.right
	for {
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.key, node.value = leftmostrightside.key, leftmostrightside.value
	node.right = remove(node.right, node.key)
	return node
}

func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()

}

func stringfy(n *Node, level int) {
	if n == nil {
		return
	}
	format := ""
	for i := 0; i < level; i++ {
		format += "    "
	}
	format += "---[ "
	level++
	stringfy(n.left, level)
	fmt.Printf(format+"%d\n", n.key)
	stringfy(n.right, level)
}
