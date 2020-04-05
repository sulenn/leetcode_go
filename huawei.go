package main

import (
	"fmt"
	"io"
)

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		fmt.Println(Solution(str))
	}
}

func Solution(str string) int {
	max := 0
	count := 0
	var word byte
	for i:=0;i<len(str);i++ {
		if word == str[i] {
			count++
		} else {
				if count > max {
					max = count
				}
				word = str[i]
				count = 1
		}
	}
	if count > max {max = count}
	return max
}
