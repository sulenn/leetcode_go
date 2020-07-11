package main

import "fmt"

//https://leetcode-cn.com/contest/weekly-contest-171/problems/convert-integer-to-the-sum-of-two-no-zero-integers/

func getNoZeroIntegers(n int) []int {
	a := 1
	b := n - a
	for !judge(a) || !judge(b) { // 任何一个数包含零均重来
		a++
		b--
	}
	return []int{a, b}
}

func judge(n int) bool { // 判断某个数有没有包含 0
	for n != 0 {
		if n%10 == 0 {
			return false
		} else {
			n = n / 10
		}
	}
	return true
}

func main() {
	fmt.Println(getNoZeroIntegers(2))
	fmt.Println(getNoZeroIntegers(11))
	fmt.Println(getNoZeroIntegers(10000))
	fmt.Println(getNoZeroIntegers(69))
	fmt.Println(getNoZeroIntegers(1010))
}
