package main

import "fmt"

//https://leetcode-cn.com/problems/ugly-number/

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		if num % 2 == 0 {
			num = num / 2
		} else if num % 3 == 0{
			num = num / 3
		} else if num % 5 == 0 {
			num = num / 5
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(14))
}
