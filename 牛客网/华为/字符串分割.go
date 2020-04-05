package main

import (
	"fmt"
	"io"
)

func main() {
	var num int
	var str string
	for {
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}
		strs := make([]string, num)
		for i:=0 ;i<num; i++ {
			fmt.Scan(&str)
			strs[i] = str
		}
		newStrs := solution(strs)
		for i:=0; i<len(newStrs); i++ {
			fmt.Println(newStrs[i])
		}
	}
}

func solution(strs []string) []string {
	newStrs := make([]string,0)
	for i:=0; i<len(strs); i++ {
		curStrs := make([]string, 0)
		for len(strs[i]) > 0 {
			var str = ""
			if len(strs[i]) >= 8 {
				str = strs[i][:8]
				strs[i] = strs[i][8:]
			} else {
				str = strs[i]
				for len(str) != 8 {
					str += "0"
				}
				strs[i]=""
			}
			curStrs = append(curStrs, str)
		}
		newStrs = append(newStrs, curStrs...)
	}
	return newStrs
}