package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/hamming-distance/

//按位异或（^）:00001001^00000101=00001100 1^1=0 0^0=0 0^1=1
//移位操作（<< >>）:<<左移 >>右移 整体左右移 不要用反
//按位与（&）: 1&1=1 1&0=0 0&0=1

//二进制解法
func hammingDistance(x int, y int) int {
	var result int
	z := x ^ y
	for z != 0 {
		if z & 1 == 1 {
			result++
		}
		z = z >> 1
	}
	return result
}

//字符串解法
//func hammingDistance(x int, y int) int {
//	xBit := intToBit(x)
//	yBit := intToBit(y)
//	result := 0
//	if len(xBit) < len(yBit) { //哪个二进制更长
//		xBit, yBit = yBit, xBit
//	}
//	for len(xBit) != len(yBit) {  // 将短的二进制补零
//		yBit = "0" + yBit
//	}
//	for i:=0 ; i< len(xBit); i++ {
//		if xBit[i] != yBit[i] {
//			result++
//		}
//	}
//	return result
//}
//
////整数转二进制
//func intToBit(x int) string {
//	if x == 0 {
//		return "0"
//	}
//	result := ""
//	for x != 0 {
//		rest := x % 2
//		result = strconv.Itoa(rest) + result
//		x = x / 2
//	}
//	return result
//}

func main() {
	fmt.Println(hammingDistance(1,4))
}
