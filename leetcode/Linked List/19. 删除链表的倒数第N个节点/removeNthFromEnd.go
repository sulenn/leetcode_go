package main

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

//一次遍历
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return nil
	}
	realHead, first, second := head, head, head // 双指针
	for i := 0; i < n; i++ {
		if first != nil {
			first = first.Next
		} else {
			return nil
		}
	}
	if first == nil {
		return second.Next
	} // 删除头节点
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return realHead
}
