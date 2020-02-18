package main

//https://leetcode-cn.com/problems/product-of-array-except-self/solution/cheng-ji-dang-qian-shu-zuo-bian-de-cheng-ji-dang-q/

import "fmt"

//乘积 = 当前数左边的乘积 * 当前数右边的乘积
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	k := 1
	for i:=0; i<len(nums); i++ {
		result[i] = k
		k *= nums[i] // 此时数组存储的是除去当前元素左边的元素乘积
	}
	k = 1
	for j:=len(nums)-1; j>=0; j-- {
		result[j] *= k // k为该数右边的乘积。
		k *= nums[j] // 此时数组等于左边的 * 该数右边的。
	}
	return result
}

func main() {
	fmt.Println(productExceptSelf([]int {1,2,3,4}))
}
