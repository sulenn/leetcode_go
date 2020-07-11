package main

import (
	"sort"
	"strconv"
)

//https://leetcode-cn.com/problems/lexicographical-numbers/
//正规的解法是10叉树 - 参考：https://leetcode-cn.com/problems/lexicographical-numbers/solution/xian-xu-bian-li-10cha-shu-by-aknifejackzhmolong/
//数字转字符串，排序后转数字
func lexicalOrder(n int) []int {
	str := make([]string, n)
	for i := 0; i < n; i++ {
		str[i] = strconv.Itoa(i + 1)
	}
	sort.Strings(str)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i], _ = strconv.Atoi(str[i])
	}
	return result
}

//func lexicalOrder(n int) []int {
//	result := make([]int, 0)
//	for i:=1; i<=9 && i<=n; i++ {
//		for j:=0; i*pow(j) <= n; j++ {
//			num := i*pow(j)
//			count := 0
//			for ; count<pow(j) && num <= n; count++ {
//				result = append(result, num)
//				num++
//			}
//		}
//	}
//	return  result
//}

func pow(num int) int {
	result := 1
	for i := 0; i < num; i++ {
		result *= 10
	}
	return result
}
