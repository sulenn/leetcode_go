package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		names := strings.Split(str,",")
		if len(names) == 0 {
			fmt.Println("error.0001")
			continue
		}
		fmt.Println(solution(names))
	}
}

func solution(names []string) string {
	nameDic := make(map[string]int)
	for _,v := range names {
		if !judge(v) {return "error.0001"}
		nameDic[v]++
	}
	minName := ""
	minNum := math.MinInt64
	for k, v := range nameDic {
		if v > minNum {
			minName = k
			minNum = v
			continue
		}
		if v == minNum && minName > k {
			minName = k
			minNum = v
		}
	}
	return minName
}

func judge(word string) bool {
	if len(word) == 0 {return false}
	if word[0] < 'A' || word[0] > 'Z' {return false}
	for i:=1;i<len(word); i++ {
		if word[i] < 'a' || word[0] > 'z' {return false}
	}
	return true
}