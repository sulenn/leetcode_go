package main

import (
	"fmt"
	"io"
)

func main() {
	var str string
	var k int
	for {
		_, err := fmt.Scan(&str, &k)
		if err == io.EOF {
			break
		}
		index := 0
		for i := 0; i < len(str); i++ {
			if str[i] < str[index] {
				index = i
			}
		}
		if k == 1 {
			fmt.Println(str[index : index+1])
		} else if k == 2 {
			if k+index-1 < len(str) {
				fmt.Println(str[index : index+2])
			}
			fmt.Println("aa")
		} else if k == 3 {
			if k+index-1 < len(str) {
				fmt.Println(str[index : index+3])
			}
			fmt.Println("aaa")
		} else if k == 4 {
			if k+index-1 < len(str) {
				fmt.Println(str[index : index+4])
			}
			fmt.Println("aaaa")
		} else if k == 5 {
			if k+index-1 < len(str) {
				fmt.Println(str[index : index+5])
			}
			fmt.Println("aaaaa")
		}
	}
}
