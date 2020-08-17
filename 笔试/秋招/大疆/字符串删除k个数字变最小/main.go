package main

import (
	"fmt"
	"io"
)

func main() {
	var str string
	var num int
	for {
		_, err := fmt.Scan(&str, &num)
		if err == io.EOF {
			break
		}
		for i := 0; i < len(str)-1; {
			if num < 1 {
				break
			}
			if str[i] > str[i+1] {
				str = str[:i] + str[i+1:]
				num--
				if i > 0 {
					i--
				}
			} else {
				i++
			}
		}
		for num > 0 && len(str) != 0 {
			str = str[:len(str)-1]
			num--
		}
		for str[0] == '0' {
			str = str[1:]
		}
		if len(str) == 0 {
			fmt.Println(0)
			continue
		}
		fmt.Println(str)
	}
}
