package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//https://www.nowcoder.com/practice/48b3cb4e3c694d9da5526e6255bb73c3?tpId=37&tqId=21236&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		strArr := strings.Split(str, " ")
		result := ""
		for i:=len(strArr)-1; i>=0; i--{
			result += " " + strArr[i]
		}
		fmt.Println(result[1:])
	}
}
