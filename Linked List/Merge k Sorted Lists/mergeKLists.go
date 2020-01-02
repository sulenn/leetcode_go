package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/merge-k-sorted-lists/

type ListNode struct {
    Val int
    Next *ListNode
}
func mergeKLists(lists []*ListNode) *ListNode {
	var result = &ListNode{0, nil}
	head := result
	for len(lists) > 0 {
		minValue := math.MaxInt64   // 每次遍历列表中最小值
		index := math.MaxInt64     // 记录最小值的下标
		length := len(lists)
		//for i:=0; i<len(lists); i++ {    // 这里条件设置为 len(lists) 和 length 均可, length 更快
		for i:=0; i<length; i++ {    // 这里条件设置为 len(lists) 和 length 均可, length 更快
			if lists[i] != nil {
				if lists[i].Val < minValue {
					minValue = lists[i].Val
					index = i
				}
			}
		}
		if index != math.MaxInt64 {  // 最小值所属的元素，链表右移一位
			head.Next = lists[index]
			head = head.Next
			lists[index] = lists[index].Next
		} else {
			break
		}
	}
	return result.Next
}

func main() {
	link1 := &ListNode{1, &ListNode{4, &ListNode{5,nil}}}
	link2 := &ListNode{1, &ListNode{3, &ListNode{4,nil}}}
	link3 := &ListNode{2, &ListNode{6, nil}}
	fmt.Println(mergeKLists([]*ListNode {link1, link2, link3}))
}
