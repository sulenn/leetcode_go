package main

//https://leetcode-cn.com/problems/sort-colors/

//快排思想，取1最为pivot，one1和one2是全1的一个区间
func sortColors(nums []int)  {
	one1 := 0 // 指向全是1的区间头
	one2 := 0 // 指向全是1的区间尾
	tail := len(nums)-1 // 指向全是2的区间头部
	for one2 <= tail {
		if nums[one2] == 1 {
			one2++
		} else if nums[one2] == 2 {  // 与末尾2区间进行交换
			nums[one2], nums[tail] = nums[tail], nums[one2]
			tail--
		} else if nums[one2] == 0 && one1 < one2 {  // 与开始0区间进行交换
			nums[one2],nums[one1] = nums[one1], nums[one2]
			one1++
			one2++
		} else {
			one2++
			one1++
		}
	}
}