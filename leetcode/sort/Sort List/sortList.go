package main

//https://leetcode-cn.com/problems/sort-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

//归并排序（从底至顶直接合并）
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	listLength := 0
	interval := 1
	lastElePerInter := head // 每一轮合并的最后一个元素，用于判断当前间隔（如间隔1，两两元素合并）回合是否结束
	for lastElePerInter != nil {
		listLength++
		lastElePerInter = lastElePerInter.Next
	}
	res := &ListNode{0, nil} // 用于定位头指针
	res.Next = head
	//以不同间隔合并列表
	for interval < listLength {
		pre := res // 每次元素比较时，指向最小的元素
		lastElePerInter = res.Next
		for lastElePerInter != nil {
			//获取每次合并的头指针，head1, head2
			head1 := lastElePerInter
			count := interval                         // 用于获取 head2 和 新的 lastElePerInter
			for count > 0 && lastElePerInter != nil { // 获取 head2
				lastElePerInter = lastElePerInter.Next
				count--
			}
			if count > 0 {
				break
			} // head2 为空，不需要合并，直接返回
			head2 := lastElePerInter
			count = interval
			for count > 0 && lastElePerInter != nil { // 获取新的 lastElePerInter
				lastElePerInter = lastElePerInter.Next
				count--
			}
			count1, count2 := interval, interval-count // 用于比较 head1 和 head2 元素
			for count1 > 0 && count2 > 0 {             // 元素两两比较
				if head1.Val < head2.Val {
					pre.Next = head1
					head1 = head1.Next
					count1--
				} else {
					pre.Next = head2
					head2 = head2.Next
					count2--
				}
				pre = pre.Next
			}
			if count1 > 0 { // 连接剩下的元素
				pre.Next = head1
			} else {
				pre.Next = head2
			}
			for count1 > 0 || count2 > 0 { // 移动指针
				pre = pre.Next
				count1--
				count2--
			}
			pre.Next = lastElePerInter // 指针连接到新的 lastElePerInter
		}
		interval *= 2 // 合并间隔加倍
	}
	return res.Next
}

func main() {

}
