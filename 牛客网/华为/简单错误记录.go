package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

//https://www.nowcoder.com/question/next?pid=260145&qid=25367&tid=31519636

//type error struct {
//	fileName string
//	line int
//	num int
//}
//
//func (this *error) judge() bool {
//	if this
//}

func main() {
	errorArr := []string{}
	errorMap := make(map[string]int)
	for {
		filePath := ""
		line := ""
		_, err := fmt.Scan(&filePath, &line)
		if err == io.EOF {
			break
		}
		if strings.Contains(filePath, "\\") { //保留文件名
			filePathArr := strings.Split(filePath, "\\")
			filePath = filePathArr[len(filePathArr)-1]
		}
		if len(filePath) > 16 { // 取文件后16位
			filePath = filePath[len(filePath)-16:]
		}
		if _, ok := errorMap[filePath+" "+line]; ok { // 累加
			errorMap[filePath+" "+line]++
		} else {
			errorArr = append(errorArr, filePath+" "+line)
			errorMap[filePath+" "+line] = 1
		}
	}
	printArr(errorArr, errorMap)
}

func printArr(errorArr []string, errorMap map[string]int) {
	for i := 0; i < 8; i++ { // 数量最多的前8个错
		max := 0
		for j := 1; j < len(errorArr); j++ {
			if errorMap[errorArr[j]] > errorMap[errorArr[max]] {
				max = j
			}
		}
		fmt.Println(errorArr[max] + " " + strconv.Itoa(errorMap[errorArr[max]]))
		errorArr = append(errorArr[:max], errorArr[max+1:]...)
	}
}
