package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l3 := &ListNode{0,nil}    //创建第三个链表过渡
	l3HeadPointer := l3    //指向链表头
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Next = l1
	} else {
		l3.Next = l2
	}
	return l3HeadPointer.Next
}

func main() {
	test := ListNode{5, nil}
	fmt.Println(test)
}
