package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//参考：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/solution/chao-qing-xi-tu-jie-san-zhi-zhen-fa-by-justdo1t/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preNode := &ListNode{0, nil}
	preNode.Next = head
	fixNode := preNode
	var firstNode *ListNode
	var secondNode *ListNode
	for fixNode.Next != nil {
		firstNode = fixNode.Next
		secondNode = fixNode.Next
		for secondNode.Next != nil && secondNode.Next.Val == firstNode.Val {
			secondNode = secondNode.Next
		}
		if firstNode == secondNode {
			fixNode = firstNode
		} else {
			fixNode.Next = secondNode.Next
		}
	}
	return preNode.Next
}

func main() {
	//fmt.Println(deleteDuplicates([]int{1,1,2,2}))
}
