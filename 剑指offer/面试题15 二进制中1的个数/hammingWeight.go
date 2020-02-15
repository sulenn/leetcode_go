package main

//https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/

//通过移位操作解题
//func hammingWeight(num uint32) int {
//	result := 0
//	for num != 0 {
//		if num&1 == 1 {
//			result++
//		}
//		num = num>>1
//	}
//	return result
//}

//更巧妙的方法，n&(n-1)将n的二进制表示中的最低位的1改为0
func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		result++
		num &= (num-1)
	}
	return result
}
