package main

import (
	"fmt"
	"math/rand"
)

//https://leetcode-cn.com/problems/implement-rand10-using-rand7/

//拒绝采样，参考：https://leetcode-cn.com/problems/implement-rand10-using-rand7/solution/chao-xiang-xi-san-chong-jie-fa-qing-song-li-jie-by/
func rand10() int {
	first := rand7()
	for first > 6 {
		first = rand7()
	}
	second := rand7()
	for second > 5 {
		first = rand7()
	}
	result := 0
	if first < 4 {
		result = 0 + second
	} else {
		result = 5 + second
	}
	return result
}

func rand7() int {
	return rand.Intn(7)
}

func main() {
	fmt.Println(rand.Intn(7))
}
