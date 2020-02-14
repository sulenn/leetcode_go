package main

import "errors"

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		panic(errors.New("数组为空！"))
	}
	if numbers[0] < numbers[len(numbers)-1] {   // 针对 [1、2] 这种特殊的测试用例
		return numbers[0]
	}
	prev := 0
	tail := len(numbers) - 1
	for tail > prev {
		mid := (prev + tail) / 2
		if mid-1 >= 0 && numbers[mid-1] > numbers[mid] {   // numbers[mid] < numbers[mid-1]
			return numbers[mid]
		} else if mid+1 < len(numbers) && numbers[mid+1] < numbers[mid] {   // numbers[mid+1] < numbers[mid]
			return numbers[mid+1]
		}
		if numbers[mid] == numbers[prev] && numbers[mid] == numbers[tail] {
			return inOrder(numbers, prev, tail)
		}
		if numbers[mid] >= numbers[prev] {
			prev = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return numbers[0]
}

func inOrder(numbers []int, prev int, tail int) int {
	for ;prev < tail; prev++ {
		if numbers[prev] > numbers[prev+1] {
			return numbers[prev+1]
		}
	}
	return numbers[prev]
}
