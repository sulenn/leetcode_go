package main

import (
	"math"
	"sort"
)

//https://leetcode-cn.com/problems/3sum-closest/

//该题与 3Sum 类似，但是更简单
func threeSumClosest(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))  // 排序
	length := len(nums)
	if length < 3 {
		return 0
	}
	result := nums[0] + nums[1] + nums[2]
	for i:=0; i<length - 2; i++ {
		left := i + 1
		right := length - 1
		for left < right {
			newValue := nums[i] + nums[left] + nums[right]
			if newValue == target {
				return target
			}
			if math.Abs(float64(newValue - target)) < math.Abs(float64(result - target)) {
				result = newValue
			}
			if newValue > target {
				right --
			} else {
				left ++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSumClosest([]int {-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int {0, 1, 2}, 3))
	//test := "qwe"
	//fmt.Println(test + "d")
}
