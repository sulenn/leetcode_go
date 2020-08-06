package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 判断有环
	slow, fast := head, head.Next
	for slow != fast {
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		slow = slow.Next
	}

	// 记录环的大小
	countNode := fast
	count := 1
	for countNode.Next != fast {
		countNode = countNode.Next
		count++
	}

	// 获得环的入口节点
	entranceNode := head
	for i := 0; i < count; i++ {
		entranceNode = entranceNode.Next
	}
	for head != entranceNode {
		head = head.Next
		entranceNode = entranceNode.Next
	}
	return entranceNode
}
