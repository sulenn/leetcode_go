package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 每一次交换，p2.next = p3.next, p3.next = p2, p1.next
// 需要注意判断临界值
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p1 *ListNode = nil
	p2 := head
	p3 := head.Next
	head = p3
	for p2 != nil && p3 != nil {
		p2.Next = p3.Next
		p3.Next = p2
		if p1 != nil { // 注意刚开始，p1 为 nil
			p1.Next = p3
		}
		p1 = p2
		p2 = p2.Next
		if p2 != nil { // 小心越界
			p3 = p2.Next
		} else {
			p3 = nil
		}
	}
	if head != nil {
		return head
	} // 链表中只有一个元素
	return p2
}
