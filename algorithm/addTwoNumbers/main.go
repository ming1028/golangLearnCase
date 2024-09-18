package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n7 := &ListNode{
		Val: 4,
	}
	n6 := &ListNode{
		Val:  6,
		Next: n7,
	}
	n3 := &ListNode{
		Val: 3,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n3,
	}

	l1 := &ListNode{
		Val:  2,
		Next: n4,
	}
	l2 := &ListNode{
		Val:  5,
		Next: n6,
	}
	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (res *ListNode) {
	carry := 0
	tail := new(ListNode)
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if res == nil {
			res = &ListNode{Val: sum}
			tail = res
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return res
}
