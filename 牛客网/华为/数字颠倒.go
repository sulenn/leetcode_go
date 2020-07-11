package main

import (
	"fmt"
	"io"
	"strconv"
)

//https://www.nowcoder.com/practice/ae809795fca34687a48b172186e3dafe?tpId=37&tqId=21234&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

func main() {
	for {
		input := 0
		_, err := fmt.Scan(&input)
		if err == io.EOF {
			break
		}
		result := ""
		if input == 0 { // 特俗情况
			fmt.Println("0")
			continue
		}
		for input != 0 {
			result += strconv.Itoa(input % 10)
			input /= 10
		}
		fmt.Println(result)
	}
}
