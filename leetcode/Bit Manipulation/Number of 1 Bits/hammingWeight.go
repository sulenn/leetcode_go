package main

//https://leetcode-cn.com/problems/number-of-1-bits/

import "fmt"

////循环和位移动
//func hammingWeight(num uint32) int {
//	result := 0
//	var bit uint32 = 1
//	for i:=0; i<32; i++ {
//		if num & bit != 0 {
//			result++
//		}
//		bit = bit << 1
//	}
//	return result
//}

//位操作的小技巧
//在二进制表示中，数字 n 中最低位的 1 总是对应 n - 1 中的 0 。因此，将 n 和 n−1 与运算总是能把 n 中最低位的 1 变成 0 ，并保持其他位不变。
func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		result++
		num = num & (num - 1)
	}
	return result
}

func main() {
	fmt.Println(hammingWeight(11))
}
