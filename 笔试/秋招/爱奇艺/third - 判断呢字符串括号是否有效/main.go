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
		queue := make([]byte, 0)
		flag := true
		for i := 0; i < len(str); i++ {
			if str[i] == '[' || str[i] == '{' || str[i] == '(' {
				queue = append(queue, str[i])
			}
			length := len(queue)
			if str[i] == '}' {
				if len(queue) > 0 && queue[length-1] == '{' {
					queue = queue[:length-1]
				} else {
					flag = false
					break
				}
			}
			if str[i] == ']' {
				if len(queue) > 0 && queue[length-1] == '[' {
					queue = queue[:length-1]
				} else {
					flag = false
					break
				}
			}
			if str[i] == ')' {
				if len(queue) > 0 && queue[length-1] == '(' {
					queue = queue[:length-1]
				} else {
					flag = false
					break
				}
			}
		}
		if len(queue) != 0 {
			flag = false
		}
		if flag {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
