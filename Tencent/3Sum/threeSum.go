package main

import "fmt"

//https://leetcode-cn.com/problems/3sum/

func threeSum(nums []int) [][]int {
	result := [][]int {}
	length := len(nums)
	nums = rank(nums)
	if length <= 2 || nums[length-1] < 0 {
		return [][]int {}
	}
	//双指针，预留两位
	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {  // 边界检测。最左边起始为整数，直接结束
			break
		}
		if i >= 1 && nums[i] == nums[i -1] {   //剪枝，去除重复的元素
			continue
		}
		pre := i + 1
		back := length - 1
		for pre < back {
			if nums[i] + nums[pre] + nums[back] == 0 {
				result = append(result, []int{nums[i], nums[pre], nums[back]})
				for pre < back && nums[pre] == nums[pre+1] {  // 去除左边指针的重复元素
					pre++
				}
				for pre < back && nums[back] == nums[back-1] {  // 去除右边指针的重复元素
					back--
				}
				pre++
				back--
			} else if nums[i] + nums[pre] + nums[back] < 0 {
				pre ++
			} else {
				back --
			}
		}
	}
	return result
}

//对数组排序（插入排序）
func rank(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j - 1] {
				nums[j], nums[j - 1] = nums[j - 1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(threeSum([]int {-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int {0, 0, 0, 0}))
	fmt.Println(threeSum([]int {-2,0,1,1,2}))
}
