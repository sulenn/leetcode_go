package main

import "strconv"

//https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

//参考：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/solution/mian-shi-ti-46-ba-shu-zi-fan-yi-cheng-zi-fu-chua-6/
//动态规划， f(i) = f(i-2) + f(i-1) 若数字 x(i-1)xi 可被翻译，否则 f(i) = f(i-1)
func translateNum(num int) int {
	str := strconv.Itoa(num)
	p1, p2 := 1, 1
	for i := 1; i < len(str); i++ {
		numStr := str[i-1 : i+1]
		if numStr >= "10" && numStr <= "25" {
			p1, p2 = p2, p1+p2
		} else {
			p1 = p2
		}
	}
	return p2
}
