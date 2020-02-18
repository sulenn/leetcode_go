package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	reverseFunc := func(digits []int) []int {
		for i := 0 ;i < len(digits)/2;i++ {
		digits[i] , digits[len(digits)-i-1] = digits[len(digits)-i-1], digits[i]
	}
		return digits
	}

	if digits[len(digits) - 1] + 1 != 10 {
		digits[len(digits) - 1]++
		return digits
	} else {
		flag := 1
		digits = reverseFunc(digits)
		for i := 0;i < len(digits);i++ {
			sum := flag + digits[i]
			if sum == 10 {
				digits[i] = 0
			} else {
				flag = 0
				digits[i] = sum
				return reverseFunc(digits)
			}
		}
		if flag == 1 {
			digits = append(digits, 1)
		}
		return reverseFunc(digits)
	}
}



func main() {
	fmt.Println(plusOne([]int{1,2,3}))
	fmt.Println(plusOne([]int{4,3,2,1}))
}
