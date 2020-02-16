package main

//https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/

type ListNode struct {
    Val int
    Next *ListNode
}

//非递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {return l2}
	if l2 == nil {return l1}
	var head  = &ListNode{0,nil}
	curNode := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curNode.Next = l1
			l1 = l1.Next
		} else {
			curNode.Next = l2
			l2 = l2.Next
		}
		curNode = curNode.Next
	}
	if l1 != nil {
		curNode.Next = l1
	} else {
		curNode.Next = l2
	}
	return head.Next
}
