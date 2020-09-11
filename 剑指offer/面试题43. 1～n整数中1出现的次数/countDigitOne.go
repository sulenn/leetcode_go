package main

import "fmt"

//https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/

//参考：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/solution/mian-shi-ti-43-1n-zheng-shu-zhong-1-chu-xian-de-2/
//1. cur=0, high×digit
//2. cur=1, high×digit+low+1
//3. cur>1, (high+1)×digit
func countDigitOne(n int) int {
	digit := 1
	high := n / 10
	low := 0
	cur := n % 10
	result := 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			result += high * digit
		}
		if cur == 1 {
			result += high*digit + low + 1
		}
		if cur > 1 {
			result += (high + 1) * digit
		}
		digit *= 10
		cur = high % 10
		high /= 10
		low = n % digit
	}
	return result
}

func main() {
	fmt.Println(countDigitOne(12))
	fmt.Println(countDigitOne(13))
}
