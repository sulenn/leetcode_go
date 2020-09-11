package main

import (
	"fmt"
	"strconv"
)

//https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/

//参考：https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/mian-shi-ti-44-shu-zi-xu-lie-zhong-mou-yi-wei-de-6/
func findNthDigit(n int) int {
	//if n == 0 {
	//	return 0
	//}
	digit := 1
	start := 1
	count := 9 * start * digit
	for n > count {
		n -= count
		digit++
		start *= 10
		count = 9 * digit * start
	}
	startNum := start + (n-1)/digit
	n = (n - 1) % digit
	result, _ := strconv.Atoi(strconv.Itoa(startNum)[n : n+1])
	return result
}

func main() {
	//fmt.Println(findNthDigit(3))
	//fmt.Println(findNthDigit(11))
	//fmt.Println(findNthDigit(10))
	//fmt.Println(findNthDigit(9))
	//fmt.Println(findNthDigit(0))
	fmt.Println("321" > "32")
}
