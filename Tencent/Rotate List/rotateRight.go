package main

//https://leetcode-cn.com/problems/rotate-list/submissions/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil { //前三个 if，排除特殊情况
		return head
	}
	if k < 0 {
		return nil
	}
	if k == 0 {
		return head
	}
	length := 1     // 计算链表长度
	curNode := head // 当前节点
	for curNode.Next != nil {
		length++
		curNode = curNode.Next
	}
	curNode.Next = head // 组成循环链表
	realK := k % length
	for i := 1; i < length-realK; i++ { // 寻找切断循环链表的点
		head = head.Next
	}
	result := head.Next // 切断循环链表
	head.Next = nil
	return result
}

func literal(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val, " -> ")
		head = head.Next
	}
}

func main() {
	temp := &ListNode{1, nil}
	head := temp
	temp.Next = &ListNode{2, nil}
	temp = temp.Next
	temp.Next = &ListNode{3, nil}
	temp = temp.Next
	temp.Next = &ListNode{4, nil}
	temp = temp.Next
	temp.Next = &ListNode{5, nil}
	temp = temp.Next
	//result := rotateRight(head, 2)
	result := rotateRight(head, 0)
	literal(result)
}
