package main

//https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/

//环形链表；超出时间限制
//func lastRemaining(n int, m int) int {
//	r := ring.New(n)
//	for i:=0; i<n; i++ {
//		r.Value = i
//		r = r.Next()
//	}
//	for r.Len()!=1 {
//		r = r.Move(m-2)  // m 是指删除从当前元素开始的第m个元素
//		r.Unlink(1)  // Unlink 是删除当前元素的下一个元素
//		r = r.Move(1)  // 删除后从删除元素的下一个开始
//	}
//	return r.Value.(int)
//}

func lastRemaining(n int, m int) int {
	result := 0
	for i := 2; i <= n; i++ {
		result = (result + m) % i
	}
	return result
}
