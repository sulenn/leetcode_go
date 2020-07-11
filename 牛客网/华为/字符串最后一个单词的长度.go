package main

import (
	"bufio"
	"fmt"
	"os"
)

//https://www.nowcoder.com/practice/8c949ea5f36f422594b306a2300315da?tpId=37&tqId=21224&tPage=1&rp=&ru=/ta/huawei&qru=/ta/huawei/question-ranking

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		result := 0
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] != ' ' {
				result++
			} else {
				break
			}
		}
		fmt.Println(result)
	}
}
