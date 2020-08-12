package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	p1, p2 := head, head
	for i := 0; i < k-1; i++ {
		if p1.Next != nil {
			p1 = p1.Next
		} else {
			return nil
		}
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}

func main() {
	var temp []int
	temp = nil
	fmt.Println(len(temp))
}
