package main

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {return l2}
	if l2 == nil {return l1}
	l1,count1 := reverseList(l1)
	l2,count2 := reverseList(l2)
	if count1 < count2 {
		l1, l2 = l2, l1
	}
	restNum := 0
	head := l1
	sum := 0
	for l1 != nil || l2 != nil {
		if l2 != nil {
			sum = l1.Val + l2.Val + restNum
			l2 = l2.Next
		} else {
			sum = l1.Val + restNum
		}
		l1.Val = sum%10
		restNum = sum/10
		l1 = l1.Next
	}
	//for l1 != nil {
	//	sum := l1.Val + restNum
	//	l1.Val = sum/10
	//	restNum = sum%10
	//	l1 = l1.Next
	//}
	if restNum != 0 {
		addNode(head, restNum)
	}
	head, _ = reverseList(head)
	return head
}

func reverseList(l *ListNode) (*ListNode, int) {
	if l == nil {return l,0}
	pre := l
	cur := pre.Next
	pre.Next = nil
	count := 1
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
		count++
	}
	return pre, count
}

func addNode(head *ListNode, val int) {
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{Val:val}
}
