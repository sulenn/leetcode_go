package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		start := -1
		newStr := ""
		for i := 0; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				if start == -1 {
					start = i
				}
			} else {
				if start != -1 {
					newStr += "*" + str[start:i] + "*"
					start = -1
				}
				newStr += str[i : i+1]
			}
		}
		if start != -1 {
			newStr += "*" + str[start:] + "*"
		}
		fmt.Println(newStr)
	}
}
