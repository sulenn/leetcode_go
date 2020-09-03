package main

import (
	"bufio"
	"fmt"
	"os"
)

//https://www.nowcoder.com/practice/d9162298cb5a437aad722fccccaae8a7?tpId=37&tqId=21227&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		inputStr := input.Text()
		for i := 0; i < len(inputStr); i += 8 {
			if len(inputStr)-i >= 8 {
				fmt.Println(inputStr[i : i+8])
			} else {
				str := inputStr[i:]
				for len(str) != 8 {
					str += "0"
				}
				fmt.Println(str)
			}
		}
	}
}
