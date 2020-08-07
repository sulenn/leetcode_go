package main

import (
	"fmt"
	"io"
)

func main() {
	var N int
	var x, y string
	for {
		_, err := fmt.Scan(&N, &x, &y)
		if err == io.EOF {
			break
		}
		fmt.Println(judge(x, y))
	}
}

func judge(str1, str2 string) int {
	num := 0
	//length := len(str2)
	for str1 != str2 {
		str2 = str2[1:] + str2[:1]
		//str2 = str2[length-1:] + str2[:length-1]
		num++
	}
	return num
}
