package main

//https://leetcode-cn.com/problems/linked-list-cycle-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	if fast.Next == nil || fast.Next.Next == nil {
		return nil
	}
	slow = slow.Next
	fast = fast.Next.Next
	for slow != fast { // 找到快慢指针的交汇点
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	intersection := slow
	for intersection != head { // 找到交汇点和头指针的交汇点，即为链表环的起点
		intersection = intersection.Next
		head = head.Next
	}
	return head
}

func main() {

}
