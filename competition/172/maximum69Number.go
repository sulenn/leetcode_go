package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/contest/weekly-contest-172/problems/maximum-69-number/

func maximum69Number (num int) int {
	length := 0
	result := 0
	flag := true
	temp_num := num
	for temp_num != 0 {
		length++
		temp_num = temp_num / 10
	}
	length--
	for ;length>=0;length--{
		if flag && num / int(math.Pow10(length)) == 6 {
			result += 9 * int(math.Pow10(length))
			num -= 6 * int(math.Pow10(length))
			flag = false
		} else {
			result += (num / int(math.Pow10(length))) * int(math.Pow10(length))
			num -= (num / int(math.Pow10(length))) * int(math.Pow10(length))
		}
	}
	return result
}

func main() {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
}
