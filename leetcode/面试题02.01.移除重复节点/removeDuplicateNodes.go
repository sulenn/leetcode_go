package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	valueDic := make(map[int]struct{})
	curNode := head
	valueDic[curNode.Val] = struct{}{}
	nextNode := curNode.Next
	for nextNode != nil {
		if _, ok := valueDic[nextNode.Val]; ok {
			curNode.Next = nextNode.Next
			nextNode = nextNode.Next
		} else {
			valueDic[nextNode.Val] = struct{}{}
			curNode = nextNode
			nextNode = nextNode.Next
		}
	}
	return head
}
