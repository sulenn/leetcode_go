package main

import (
	"bufio"
	"fmt"
	"os"
)

//https://www.nowcoder.com/practice/253986e66d114d378ae8de2e6c4577c1?tpId=37&tqId=21232&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		num := input.Text()
		flagArr := make([]bool, 10)   // 0 - 10
		result := ""
		for i:=len(num)-1; i>=0; i-- {
			if !flagArr[num[i]-'0'] {
				result += string(num[i])
				flagArr[num[i]-'0'] = true
			}
		}
		fmt.Println(result)
	}
}
