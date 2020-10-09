package main

//https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/

//思路：将数组A排序，获取数组中最大的数max，申请与max同样大小的数组arr，遍历A，数组arr记录每个数字出现的次数。
//用两个指针，分别指向出现次数大于等于2的数字和一次都没有出现的数字，累计相减，差值累加。
func minIncrementForUnique(A []int) int {
	max := 0
	for _, v := range A {
		if v > max {
			max = v
		}
	}
	arr := make([]int, max+1) // 获取A数组中每个数字出现的次数，注意长度为max+1
	for _, v := range A {
		arr[v]++
	}
	numP := 0   // 指向有数字的指针
	noneP := 0  // 指向没有数字的指针
	result := 0 // 移动总次数
	for numP < max+1 {
		for numP < max+1 && (arr[numP] == 0 || arr[numP] == 1) { // 频率大于等于2的数字
			numP++
		}
		for noneP < max+1 && arr[noneP] != 0 || noneP < numP { // 频率为0的数字
			noneP++
		}
		if numP >= max+1 {
			break
		}
		result += noneP - numP
		noneP++
		arr[numP]--
	}
	return result
}
