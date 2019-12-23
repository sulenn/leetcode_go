package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/gray-code/

func grayCode(n int) []int {
	if n == 0 {  //特殊情况
		return []int {0}
	}
	if n == 1 {   //递归结束
		return []int {0, 1}
	}
	result := []int {}
	recursiveResult := []int {}
	if n > 1 {   // 递归开始
		recursiveResult = grayCode(n-1)
	}
	result = append(result, recursiveResult...)   //正向情况，复制
	recursiveResult = reverse(recursiveResult)    //反向，反转
	for i:=0; i<len(recursiveResult); i++ {      //二进制加 1，就是每个数加上 2 的 n-1 幂
		result = append(result, recursiveResult[i] + int(math.Pow(2, float64(n-1))))
	}
	return result
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(3))
}
