package main

import (
	"fmt"
	"io"
)

//https://www.nowcoder.com/practice/bd891093881d4ddf9e56e7cc8416562d?tpId=182&&tqId=34785&rp=1&ru=/activity/oj&qru=/ta/exam-all/question-ranking

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		max := ""
		for i := 0; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				length := 1
				index := i + 1
				for index < len(str) && str[index] >= '0' && str[index] <= '9' {
					length++
					index++
				}
				if length > len(max) {
					max = str[i:index]
				}
				i = index
			}
		}
		fmt.Println(max)
	}
}
