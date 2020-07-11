package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//先反转链表
func reversePrint(head *ListNode) []int {
	length := 0
	var preNode *ListNode = nil //  这么不能用 preNode := new(ListNode),new(ListNode) 的值为{0,nil}
	curNode := head
	for curNode != nil { // 反转链表
		temp := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = temp
		length++
	}
	result := make([]int, length)
	i := 0
	for preNode != nil { // 输出反转后的链表节点值
		result[i] = preNode.Val
		preNode = preNode.Next
		i++
	}
	return result
}

//获取链表长度，声名数组，遍历链表至数组，翻转数组
//func reversePrint(head *ListNode) []int {
//	length := 0
//	tempPointer := head
//	for tempPointer != nil {  // 获取链表长度
//		length++
//		tempPointer = tempPointer.Next
//	}
//	result := make([]int, length)
//	i:=0
//	for head != nil {  // 遍历链表至数组
//		result[i] = head.Val
//		head = head.Next
//		i++
//	}
//	for j:=0; j<length/2; j++ {  // 翻转数组
//		result[j], result[length-j-1] = result[length-j-1], result[j]
//	}
//	return result
//}
