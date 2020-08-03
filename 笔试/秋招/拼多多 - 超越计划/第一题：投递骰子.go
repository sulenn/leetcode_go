package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var target int
	var count int
	for {
		_, err := fmt.Scan(&target, &count)
		if err == io.EOF {
			break
		}
		nums := make([]int, count)
		for i := 0; i < count; i++ {
			fmt.Scan(&nums[i])
		}
		sum := 0
		flag := false
		backTimes := 0
		if target == 0 {
			fmt.Println("paradox")
			continue
		}
		for i := 0; i < count; i++ {
			nums[i] %= target * 2
			backTimes += nums[i] / (target * 2)
			if sum+nums[i] == target {
				fmt.Println("paradox")
				flag = true
				break
			}
			if sum+nums[i] > target {
				sum = 2*target - nums[i] - sum
				backTimes++
			} else {
				sum += nums[i]
			}
		}
		if flag {
			continue
		}
		fmt.Println(strconv.Itoa(target-sum) + " " + strconv.Itoa(backTimes))
	}
}
