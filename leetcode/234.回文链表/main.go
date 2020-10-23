package main

import "fmt"

//https://leetcode-cn.com/problems/palindrome-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

//参考：https://leetcode-cn.com/problems/palindrome-linked-list/solution/cjian-ji-dai-ma-by-orangeman-27/
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	var pre *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// reverse slow list
		temp := slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}
	if fast != nil {
		slow = slow.Next
	}
	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}

func main() {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 2, Next: nil}
	node4 := &ListNode{Val: 1, Next: nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	fmt.Println(isPalindrome(node1))
}
