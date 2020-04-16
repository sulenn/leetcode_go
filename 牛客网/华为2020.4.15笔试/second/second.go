package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var str1 string
	var str2 string
	for {
		_, err := fmt.Scan(&str1, &str2)
		if err == io.EOF {
			break
		}
		fmt.Print(solution(str1,str2))
	}
}

func solution(str1, str2 string) string {
	str1 += "[addr="
	strs := strings.Split(str2, str1)
	result := ""
	for i:=1; i<len(strs);i++ {
		subStr := strings.Split(strs[i],",")
		addr := subStr[0]
		if !judge(addr) {return "FAIL"}
		mask := subStr[1][5:]
		if !judge(mask) {return "FAIL"}
		val := subStr[2][4:len(subStr[2])-1]
		if !judge(val) {return "FAIL"}
		result += addr + " " + mask + " " + val + "\r\n"
	}
	if len(result) == 0 {return "FAIL"}
	return result
}

func judge(word string) bool {
	if word[0] != '0' {return false}
	if word[1] == 'x' || word[1] == 'X' {
		return true
	} else {
		return false
	}
}
