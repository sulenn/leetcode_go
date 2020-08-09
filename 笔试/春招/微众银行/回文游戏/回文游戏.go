package main

import (
	"fmt"
	"io"
)

func main() {
	var nums int
	var str string
	for {
		_, err := fmt.Scan(&nums)
		if err == io.EOF {
			break
		}
		for i := 0; i < nums; i++ {
			fmt.Scan(&str)
			fmt.Println(solution(str))
		}
	}
}

func solution(str string) string {
	dic := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		dic[str[i]]++
	}
	count := 0
	for _, v := range dic {
		if v%2 == 1 {
			count++
		}
	}
	if count == 0 || count%2 == 1 {
		return "Cassidy"
	}
	return "Eleanore"
}
