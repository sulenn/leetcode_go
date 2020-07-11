package main

import "fmt"

//https://leetcode-cn.com/problems/count-numbers-with-unique-digits/

//写下前几个，就能看出规律了。
//
//n=1: res=10
//
//n=2: res=9*9+10=91 # 两位数第一位只能为1-9，第二位只能为非第一位的数，加上一位数的所有结果
//
//n=3: res=9 * 9 * 8+91=739 # 三位数第一位只能为1-9，第二位只能为非第一位的数，第三位只能为非前两位的数，加上两位数的所有结果
//
//n=4: res=9 * 9 * 8 * 7+739=5275 # 同上推法

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	result := 10
	cul := 9                           // 最高位不能为 0，所以9种可能
	for i := 1; i < n && i < 10; i++ { // 第二位可以为 0，也有9种可能。第三位8种可能
		cul *= (10 - i)
		result += cul
	}
	return result
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(2))
	fmt.Println(countNumbersWithUniqueDigits(1))
	fmt.Println(countNumbersWithUniqueDigits(0))
	fmt.Println(countNumbersWithUniqueDigits(3))
}
