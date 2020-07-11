package main

import "math"

//https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/

//该解法没有考虑到大数的情况，不是 “剑指 offer” 想表达的意思
func printNumbers(n int) []int {
	maxNumer := int(math.Pow10(n) - 1)
	result := make([]int, maxNumer)
	for i := 0; i < maxNumer; i++ {
		result[i] = i + 1
	}
	return result
}
