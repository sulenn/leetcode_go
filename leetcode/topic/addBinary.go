package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) < len(b) { //比较，找出最长的字符串
		a, b = b, a
	}
	var flag uint8 = 0 //标志位，是否进一
	index := 0         //字符串索引下标
	newStr := ""       //新的字符串
	aLen := len(a)
	bLen := len(b)
	for ; index < aLen; index++ {
		var sum uint8 = 0
		if bLen > index { //判断 b 字符串是否已经到达边界
			sum = flag + a[aLen-index-1] + b[bLen-index-1] - 48 - 48 // 1 对应的 asc ii 码为 49，所以这儿是减 48
		} else {
			sum = flag + a[aLen-index-1] - 48
		}
		if sum > 1 {
			newStr = string(sum-2+48) + newStr
			flag = 1
		} else {
			newStr = string(sum+48) + newStr
			flag = 0
		}
	}
	if flag == 1 {
		newStr = "1" + newStr
	}
	return newStr
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
