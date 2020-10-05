package main

import (
	"fmt"
	"io"
)

func main() {
	var N, M int
	for {
		_, err := fmt.Scan(&N, &M)
		if err == io.EOF {
			break
		}
		if M == 1 {
			fmt.Println(1)
		} else if M == 2 {
			if N == 1 {
				fmt.Println(1)
				continue
			}
			num1 := 1
			num2 := 2
			for i := 3; i <= N; i++ {
				num2, num1 = num1+num2, num2
				num2 %= 1000000007
			}
			fmt.Println(num2)
		} else if M == 3 {
			if N == 1 {
				fmt.Println(1)
				continue
			} else if N == 2 {
				fmt.Println(3)
				continue
			}
			num1 := 1
			num2 := 3
			for i := 3; i <= N; i++ {
				num2, num1 = num1*2+num2, num2
				num2 %= 1000000007
			}
			fmt.Println(num2)
		} else if M == 4 {
			if N == 1 {
				fmt.Println(1)
				continue
			} else if N == 2 {
				fmt.Println(5)
				continue
			}
			num1 := 1
			num2 := 5
			for i := 3; i <= N; i++ {
				num2, num1 = num1*5+num2, num2
				num2 %= 1000000007
			}
			fmt.Println(num2)
		} else if M == 5 {
			if N == 1 {
				fmt.Println(1)
				continue
			} else if N == 2 {
				fmt.Println(8)
				continue
			}
			num1 := 1
			num2 := 8
			for i := 3; i <= N; i++ {
				num2, num1 = num1*8+num2, num2
				num2 %= 1000000007
			}
			fmt.Println(num2)
		}
	}
}
