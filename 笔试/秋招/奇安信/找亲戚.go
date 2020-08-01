package main

import "fmt"

/**
 * 返回亲7数个数
 * @param digit int整型一维数组 组成亲7数的数字数组
 * @return int整型
 */
func reletive_7(digit []int) int {
	count := 0
	backtracking(digit, []int{}, &count)
	return count
}

func backtracking(digit []int, nums []int, count *int) {
	if len(digit) == 0 {
		if intListToInt(nums)%7 == 0 {
			*count = *count + 1
		}
		return
	}
	for i := 0; i < len(digit); i++ {
		newDigit := make([]int, len(digit)-1)
		copy(newDigit[:i], digit[:i])
		copy(newDigit[i:], digit[i+1:])
		backtracking(newDigit, append(nums, digit[i]), count)
	}
}

func intListToInt(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result*10 + nums[i]
	}
	return result
}

func main() {
	fmt.Println(reletive_7([]int{1, 1, 2, 2}))
}
