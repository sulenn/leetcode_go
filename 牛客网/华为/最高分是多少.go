package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		stuNum, doNum := 0, 0
		_, err := fmt.Scan(&stuNum, &doNum)  // 获取学生数量和操作次数
		if err == io.EOF {
			break
		}
		stuArr := make([]int, stuNum)
		for i:=0 ;i<stuNum; i++ { // 循环获取学生成绩
			_, err = fmt.Scan(&stuArr[i])
		}
		for i:=0; i<doNum; i++ {
			do := ""
			A, B := 0, 0
			_, err = fmt.Scan(&do, &A, &B)  // 获取每次操作
			if do == "Q" {
				if A > B {
					A,B = B,A
				}
				max := 0
				for ;A<=B; A++ {
					if stuArr[A-1] > max {
						max = stuArr[A-1]
					}
				}
				fmt.Println(max)
			} else {
				stuArr[A-1] = B
			}
		}
	}
}
