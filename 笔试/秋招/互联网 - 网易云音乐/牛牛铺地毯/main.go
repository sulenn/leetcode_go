package main

import (
	"fmt"
	"io"
)

func main() {
	var T int
	for {
		_, err := fmt.Scan(&T)
		if err == io.EOF {
			break
		}
		var num int
		for i := 0; i < T; i++ {
			fmt.Scan(&num)
			fmt.Println(calculate(num))
		}
	}
}

func calculate(num int) int {
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 2
	}
	if num == 3 {
		return 4
	}
	num1, num2, num3 := 1, 2, 4
	for i := 4; i <= num; i++ {
		num1, num2, num3 = num2, num3, (num1+num2+num3)%10007
	}
	return num3
}
