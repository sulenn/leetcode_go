package main

import "fmt"

// 可以参考 Bit Manipulation/Number of Bits/hammingweight.go
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	result := []int{0}
	for i := 1; i <= num; i++ {
		//在二进制表示中，数字 n 中最低位的 1 总是对应 n - 1 中的 0 。因此，将 n 和 n−1 与运算总是能把 n 中最低位的 1 变成 0 ，并保持其他位不变。
		result = append(result, result[i&(i-1)]+1) // 位操作小技巧
	}
	return result
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
