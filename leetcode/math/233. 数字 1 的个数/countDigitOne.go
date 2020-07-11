package main

//https://leetcode-cn.com/problems/number-of-digit-one/

//参考：https://leetcode-cn.com/problems/number-of-digit-one/solution/shu-zi-1-de-ge-shu-by-leetcode/
//每一位上1的数量公式：countr += (n / divider) * i + min(max(n % divider - i + 1, 0), i);
func countDigitOne(n int) int {
	result := 0
	for i := 1; i <= n; i *= 10 {
		divider := i * 10
		temp1 := 0
		temp2 := 0
		if n%divider-i+1 > 0 {
			temp1 = n%divider - i + 1
		} else {
			temp1 = 0
		}
		if temp1 < i {
			temp2 = temp1
		} else {
			temp2 = i
		}
		result += n/divider*i + temp2
	}
	return result
}
