package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 参考：https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 判断有环
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}
