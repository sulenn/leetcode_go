package main

import (
	"fmt"
	"io"
)

//https://www.nowcoder.com/practice/a35ce98431874e3a820dbe4b2d0508b1?tpId=37&tqId=21225&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

func main() {
	for {
		line1 := ""
		line2 := ""
		_, err := fmt.Scan(&line1, &line2)
		if err == io.EOF {
			break
		}
		result := 0
		for i:=0; i<len(line1); i++ {
			if line1[i] == line2[0] || line1[i]+'A'-'a' == line2[0] || line1[i]+'a'-'A' == line2[0] {
				result++
			}
		}
		fmt.Println(result)
	}
}
