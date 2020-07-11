package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		input := ""
		_, err := fmt.Scan(&input)
		if err == io.EOF {
			break
		}
		dic := map[byte]int{'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15}
		result := 0
		for i := len(input) - 1; i >= 2; i-- {
			if _, ok := dic[input[i]]; ok {
				result += dic[input[i]] * pow(len(input)-i-1)
			} else {
				result += int(input[i]-'0') * pow(len(input)-i-1)
			}
		}
		fmt.Println(result)
	}
}

func pow(num1 int) int {
	result := 1
	for i := 0; i < num1; i++ {
		result *= 16
	}
	return result
}
