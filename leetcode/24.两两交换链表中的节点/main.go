package main

//https://leetcode-cn.com/problems/swap-nodes-in-pairs/

type ListNode struct {
    Val int
    Next *ListNode
}

//参考：https://leetcode-cn.com/problems/swap-nodes-in-pairs/solution/
func swapPairs(head *ListNode) *ListNode {
	var preNode *ListNode = &ListNode{Val: 0, Next: head}
	temp := preNode
	for temp.Next != nil && temp.Next.Next != nil {
		start := temp.Next
		end := temp.Next.Next
		temp.Next = end
		start.Next = end.Next
		end.Next = start
		temp = start
	}
	return preNode.Next
}