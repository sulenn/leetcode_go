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
		fillSuffix := ""
		for i := 0; i < len(str); i++ {
			if judgeStr(str + fillSuffix) {
				fmt.Println(str + fillSuffix)
				break
			} else {
				fillSuffix = string(str[i]) + fillSuffix
			}
		}
	}
}

func judgeStr(str string) bool {
	head := 0
	tail := len(str) - 1
	for head < tail {
		if str[head] != str[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}
