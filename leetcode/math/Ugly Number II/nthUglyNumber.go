package main

import "fmt"

//https://leetcode-cn.com/problems/ugly-number-ii/

//动态规划的方法
//uglyAyrray[i] 表示第i个丑数
//uglyAyrray[i] = min(2 * uglyAyrray[l2], 3 * uglyAyrray[l3], 5 * uglyAyrray[l5])
//这里 l2, l3, l5是表示两倍、三倍、五倍时uglyAyrray列表的下标指到的位置
//参考：https://leetcode-cn.com/problems/ugly-number-ii/solution/dui-he-dong-tai-gui-hua-by-powcai/
func nthUglyNumber(n int) int {
	uglyAyrray := make([]int, n)
	uglyAyrray[0] = 1
	l2 := 0
	l3 := 0
	l5 := 0
	var min func(l2 int, l3 int, l5 int) int
	min = func(l2 int, l3 int, l5 int) int {   // 获取三个抽数中最小值
		if 2*uglyAyrray[l2] < 3*uglyAyrray[l3] {
			if 2*uglyAyrray[l2] < 5*uglyAyrray[l5] {
				return 2*uglyAyrray[l2]
			} else {
				return 5*uglyAyrray[l5]
			}
		} else {
			if 3*uglyAyrray[l3] < 5*uglyAyrray[l5] {
				return 3*uglyAyrray[l3]
			} else {
				return 5*uglyAyrray[l5]
			}
		}
	}
	for i:=1; i<n; i++ {
		uglyAyrray[i] = min(l2, l3, l5)
		if uglyAyrray[i] == 2*uglyAyrray[l2] {l2++}  //三个 if 用于排除重复元素
		if uglyAyrray[i] == 3*uglyAyrray[l3] {l3++}
		if uglyAyrray[i] == 5*uglyAyrray[l5] {l5++}
	}
	return uglyAyrray[n-1]
}


//使用 Hash map 的方法，超时
//func nthUglyNumber(n int) int {
//	uglyMap := make(map[int]bool)
//	result := 0
//	if n <= 0 { return 0 }
//	if n == 1 { return 1 }
//	uglyMap[1]=true
//	n--
//	count := 2
//	for n > 0 {
//		if count%2==0 {
//			if _, ok := uglyMap[count / 2]; ok{
//				uglyMap[count] = true
//				result = count
//				n--
//			}
//		} else if count%3==0 {
//			if _, ok := uglyMap[count / 3]; ok {
//				uglyMap[count] = true
//				result = count
//				n--
//			}
//		} else if count%5==0 {
//			if _, ok := uglyMap[count / 5]; ok {
//				uglyMap[count] = true
//				result = count
//				n--
//			}
//		}
//		count++
//	}
//	return result
//}

func main() {
	fmt.Println(nthUglyNumber(10))
	//fmt.Println(nthUglyNumber(11))
}
