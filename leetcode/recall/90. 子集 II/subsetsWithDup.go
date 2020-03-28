package main

import "sort"

//https://leetcode-cn.com/problems/subsets-ii/description/

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int {{}}
	start := 0
	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {   // 累计重复元素的步长
			 continue
		} else {
			num := len(result)   // 添加当前元素之前已有的数组长度
			for j:=i ; j>=start; j-- {
				for _,v := range result[len(result)-num:] {
					// 注意不能用这种方式，会出错。比如有个值为{1,2,3}，length:=3,capacity=4.由这个值可以生成{1,2,3,4}。一旦之后出现{1,2,3,5}
					//就会将之前的{1,2,3,4}改变为{1,2,3,5}
					//result = append(result, append(v,nums[i]))
					result = append(result, append([]int{nums[i]},v...))     // 特别注意这里，不能用append(v,nums[i])。会出现覆盖导致错误
				}
			}
			start = i+1
		}
	}
	num := len(result)  // 剩下最后一个元素
	for j:=len(nums) ; j>start; j-- {
		for _,v := range result[len(result)-num:] {
			result = append(result, append([]int{nums[len(nums)-1]},v...))
		}
	}
	return result
}
