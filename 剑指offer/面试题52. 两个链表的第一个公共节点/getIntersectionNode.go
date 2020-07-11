package 面试题52__两个链表的第一个公共节点

//https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	newHeadA := headA
	newHeadB := headB

	for newHeadA != newHeadB {
		if newHeadA == nil {
			newHeadA = headB
		} else {
			newHeadA = newHeadA.Next
		}
		if newHeadB == nil {
			newHeadB = headA
		} else {
			newHeadB = newHeadB.Next
		}
	}
	return newHeadA
}
