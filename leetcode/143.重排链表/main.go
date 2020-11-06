package main

//https://leetcode-cn.com/problems/reorder-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	nodeDic := make(map[int]*ListNode, 0)
	length := 0
	node := head
	for node != nil {
		nodeDic[length] = node
		length++
		node = node.Next
	}
	node = head
	for i := 1; i < (length+1)/2; i++ {
		temp := nodeDic[length-i]
		temp.Next = node.Next
		node.Next = temp
		node = temp.Next
		temp = nodeDic[length-i-1] // left node of insertedNode needs to remove Next
		temp.Next = nil
	}
}

func printNode(node *ListNode) {
	for node != nil {
		println(node.Val)
		node = node.Next
	}
}

func main() {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	printNode(node1)
	reorderList(node1)
	//printNode(node1)
}
