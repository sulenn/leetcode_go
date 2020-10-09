package main

//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	p := &ListNode{0, nil} // 指向头指针
	t := p                 // 指向尾指针
	curNode := head        // 用于计算链表剩余结点数，判断是否满足一轮 k
	count := 0
	for curNode != nil {
		count = 0
		for curNode != nil && count < k { // 判断剩余节点数量是否满足 k
			curNode = curNode.Next
			count++
		}
		if count == k {
			curTail, curHead, restHead := reverseByK(head, k)
			t.Next = curHead
			t = curTail
			head = restHead
		} else { // 剩余节点数量不满足 k
			t.Next = head
		}
	}
	return p.Next
}

// 反转 k 个几点，返回反转后的链表和剩余的链表
func reverseByK(head *ListNode, k int) (*ListNode, *ListNode, *ListNode) { // 分别为分离出来的子链表尾、子链表头和剩余链表头
	if head == nil {
		return nil, nil, nil
	}
	first := head
	second := first.Next
	first.Next = nil
	count := 1
	for second != nil && count < k {
		temp := second.Next
		second.Next = first
		first = second
		second = temp
		count++
	}
	return head, first, second
}
