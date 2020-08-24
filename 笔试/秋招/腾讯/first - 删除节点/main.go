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
		str := input.Text()
		kStr := str[len(str)-1:]
		k, _ := strconv.Atoi(kStr)
		input.Scan()
		str = input.Text()
		//if err == io.EOF {
		//	break
		//}

		//li := make([]int, n+1)
		//for i := 1; i <= n; i++ {
		//	fmt.Scan(&li[i])
		//}
		//input := bufio.NewScanner(os.Stdin)
		//input.Scan()
		//str := input.Text()

		realIndex := (k - 1) * 2
		fmt.Println(str[:realIndex] + str[realIndex+2:])
	}
}
