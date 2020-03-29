package main

import (
	"fmt"
	"io"
)

func main() {
	var str string
	for {
		_, err := fmt.Scanln(&str)
		if err == io.EOF {
			break
		}
		fmt.Println(solution1(str))
	}
}

func solution1(str string) int {
	if len(str) == 0{return 0}
	arr := []rune(str)
	result :=0
	start := 0
	dic := make(map[rune]int)
	for i:=0; i<len(arr); i++ {
		if arr[i] != arr[start]  {
			if i-start > dic[arr[start]] {
				dic[arr[start]] = i-start
			}
			start = i
		}
	}
	if len(arr)-start > dic[arr[start]] {
		dic[arr[start]] = len(arr)-start
	}
	for _, v := range dic {
		result += v
	}
	return result
}