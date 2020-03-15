package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		num, _ := strconv.Atoi(line)
		result := []int {}
		for i:=2; i<= num; i++ {   // 这里优化的话可以用 sqrt(num) 取根号
			for num % i == 0 {
				result = append(result, i)
				num = num / i
			}
		}
		str := ""
		for _, v:= range result {
			str += strconv.Itoa(v) + " "
		}
		fmt.Println(str)
	}
}