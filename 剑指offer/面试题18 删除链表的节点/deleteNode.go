package main

type ListNode struct {
    Val int
    Next *ListNode
}
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	tempHead := head
	if tempHead.Val == val {
		return tempHead.Next
	}
	for tempHead.Next != nil {
		if tempHead.Next.Val == val {
			tempHead.Next = tempHead.Next.Next
			break
		}
		tempHead = tempHead.Next
	}
	return head
}
