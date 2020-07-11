package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.nowcoder.com/practice/3ab09737afb645cc82c35d56a5ce802a?tpId=37&tqId=21230&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

//以小数点为分割线，分为左和右
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		num1, _ := strconv.Atoi(strings.Split(line, ".")[0])
		num2 := strings.Split(line, ".")[1]
		if num2[0] >= '5' {
			fmt.Println(num1 + 1)
		} else {
			fmt.Println(num1)
		}
	}
}
