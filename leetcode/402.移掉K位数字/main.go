package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/remove-k-digits/

func removeKdigits(num string, k int) string {
	i := 0
	count := 0
	numArr := []byte(num)
	newNumArr := make([]byte, 0)
	for ; i < len(numArr) && count < k; i++ {
		if i+1 < len(numArr) && numArr[i] > numArr[i+1] {
			count++
			for len(newNumArr) > 0 && count < k { // 模拟弹栈
				if newNumArr[len(newNumArr)-1] > numArr[i+1] {
					newNumArr = newNumArr[:len(newNumArr)-1]
					count++
				} else {
					break
				}
			}
			continue
		}
		newNumArr = append(newNumArr, numArr[i])
	}
	newNumArr = append(newNumArr, numArr[i:]...)
	for count < k && len(newNumArr) > 0 {
		newNumArr = newNumArr[:len(newNumArr)-1]
		count++
	}
	result := string(newNumArr)
	result = strings.TrimLeft(result, "0")
	if len(result) == 0 {
		return "0"
	}
	return result
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
	fmt.Println(removeKdigits("1234567890", 9))
}
