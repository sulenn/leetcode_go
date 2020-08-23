package main

import (
	"fmt"
	"io"
	"strconv"
)

var result [][]int

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		result = make([][]int, 0)
		backtracking([]int{}, 3, n)
		fmt.Println(len(result))
		for i := 0; i < len(result); i++ {
			fmt.Println(strconv.Itoa(result[i][0]) + " " + strconv.Itoa(result[i][1]))
		}
	}
}

func backtracking(nums []int, num int, target int) {
	if num == 0 {
		if nums[0] == nums[1] || nums[0] == nums[2] || nums[1] == nums[2] {
			return
		}
		num1 := nums[0]*100 + nums[1]*10 + nums[2]
		num2 := nums[0]*100 + nums[2]*10 + nums[2]
		if num1+num2 == target && num1 != num2 {
			result = append(result, []int{num1, num2})
		}
		return
	}
	var index int
	if num == 3 {
		index = 1
	} else {
		index = 0
	}
	for ; index < 10; index++ {
		newNums := make([]int, len(nums)+1)
		copy(newNums, nums)
		newNums[len(nums)] = index
		backtracking(newNums, num-1, target)
	}
}
