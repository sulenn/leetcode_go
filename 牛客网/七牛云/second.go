package main

import (
	"fmt"
	"strconv"
)

/**
 *
 * @param input string字符串一维数组 逆波兰表达式
 * @return int整型
 */
func ReversePolishNotation1( input []string ) int {
	// write code here
	if len(input) == 0 {return 0}
	arr := []int {}
	for i:=0; i< len(input); i++ {
		if input[i] == "+" {
			if len(arr) < 2 {return -1}
			arr[len(arr)-2] = arr[len(arr)-1]+arr[len(arr)-2]
			arr = arr[:len(arr)-1]
		} else if input[i] == "-" {
			if len(arr) < 2 {return -1}
			arr[len(arr)-2] = arr[len(arr)-2]-arr[len(arr)-1]
			arr = arr[:len(arr)-1]
		} else if input[i] == "*" {
			if len(arr) < 2 {return -1}
			arr[len(arr)-2] = arr[len(arr)-1]*arr[len(arr)-2]
			arr = arr[:len(arr)-1]
		} else if input[i] == "/" {
			if len(arr) < 2 {return -1}
			arr[len(arr)-2] = arr[len(arr)-2]/arr[len(arr)-1]
			arr = arr[:len(arr)-1]
		} else {
			num, _ := strconv.Atoi(input[i])
			arr = append(arr, num)
		}
	}
	return arr[0]
}

func main() {
	fmt.Println(ReversePolishNotation1([]string{"3","4","5","*","1","+","-"}))
}