package main

import (
	"fmt"
	"io"
)

//https://www.nowcoder.com/practice/944e5ca0ea88471fbfa73061ebe95728?tpId=182&&tqId=34528&rp=1&ru=/activity/oj&qru=/ta/exam-all/question-ranking

func main() {
	var N int
	for {
		_, err := fmt.Scan(&N)
		if err == io.EOF {
			break
		}
		sum := 0
		value := 1024 - N
		sum += value / 64
		value = value % 64
		sum += value / 16
		value = value % 16
		sum += value / 4
		value = value % 4
		sum += value
		fmt.Println(sum)
	}
}
