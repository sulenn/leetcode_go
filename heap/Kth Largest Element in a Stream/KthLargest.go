package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/

type KthLargest struct {
	nums []int
	k int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{nums:nums, k:k}
}

func (this *KthLargest) Add(val int) int {
	// 两种特殊情况：1.初始切片为空；2.插入数字大于切片中最大的数字
	if len(this.nums)==0 || val >= this.nums[len(this.nums)-1] {
		this.nums = append(this.nums, val)
		return this.nums[len(this.nums)-this.k]
	}
	// 正常插入
	for i:=0; i<len(this.nums); i++ {
		if val < this.nums[i] {
			temp := append([]int {}, this.nums[i:]...)
			this.nums = append(this.nums[0:i],val)
			this.nums = append(this.nums[0:i+1], temp...)
			break
		}
	}
	return this.nums[len(this.nums)-this.k]
}

func main() {
	kth := Constructor(3, []int {4,5,8,2})
	fmt.Println(kth.Add(3))
	fmt.Println(kth.Add(5))
	fmt.Println(kth.Add(10))
	fmt.Println(kth.Add(9))
	fmt.Println(kth.Add(4))
}
