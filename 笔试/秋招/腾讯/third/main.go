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
		var value int
		for i := 0; i < T; i++ {
			fmt.Scanln(&value)
			num1, num2 := one(value)
			fmt.Println(resul(num1, num2))
		}
	}
}

func one(num int) (int, int) {
	var num1, num2 int
	num1 = 9
	for num1*10+9 <= num {
		num1 = 10*num1 + 9
	}
	num2 = num - num1
	return num1, num2
}

func resul(num1, num2 int) int {
	result := 0
	for num1 != 0 {
		result += num1 % 10
		num1 /= 10
	}
	for num2 != 0 {
		result += num2 % 10
		num2 /= 10
	}
	return result
}
