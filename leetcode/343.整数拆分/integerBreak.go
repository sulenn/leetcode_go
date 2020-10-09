package _43__整数拆分

//https://leetcode-cn.com/problems/integer-break/

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	result := 1
	for n > 4 {
		result *= 3
		n -= 3
	}
	return result * n
}
