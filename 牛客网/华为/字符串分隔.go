package main

import (
	"fmt"
	"io"
)

//https://www.nowcoder.com/practice/d9162298cb5a437aad722fccccaae8a7?tpId=37&tqId=21227&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

func main() {
	for {
		input := ""
		_, err := fmt.Scan(&input)
		if err == io.EOF {
			break
		}
		for i:=0; i<len(input); i+=8 {
			if len(input) - i >= 8 {
				fmt.Println(input[i:i+8])
			} else {
				str := input[i:]
				for len(str) != 8 {
					str += "0"
				}
				fmt.Println(str)
			}
		}
	}
}
