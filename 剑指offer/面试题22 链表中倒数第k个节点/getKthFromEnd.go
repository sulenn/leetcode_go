package main

//https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

type ListNode struct {
    Val int
    Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {return nil}  //注意鲁棒
	firstPoint, secondPoint := head, head   //双指针
	for i:=1; i<k; i++ {
		if firstPoint.Next != nil {
			firstPoint = firstPoint.Next
		} else {
			return nil
		}
	}
	for firstPoint.Next != nil {
		firstPoint = firstPoint.Next
		secondPoint = secondPoint.Next
	}
	return secondPoint
}
