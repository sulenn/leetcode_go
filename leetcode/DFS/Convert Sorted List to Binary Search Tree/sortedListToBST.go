package main

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val:head.Val}
	}
	if head.Next.Next == nil {
		leftTree := TreeNode{Val:head.Val}
		return &TreeNode{head.Next.Val, &leftTree, nil}
	}
	pre := head
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	fast = slow.Next
	root := TreeNode{Val:slow.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(fast)
	return &root
}

func main() {
	
}
