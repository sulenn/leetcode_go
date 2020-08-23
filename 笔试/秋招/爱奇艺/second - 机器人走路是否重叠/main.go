package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		pathDic := make(map[string]struct{})
		x, y := 0, 0
		//temp := strconv.Itoa(x) + "++++" + strconv.Itoa(y)
		//pathDic[temp] = struct{}{}
		flag := true
		for i := 0; i < len(str); i++ {
			if str[i] == 'E' {
				y++
			}
			if str[i] == 'S' {
				x--
			}
			if str[i] == 'W' {
				y--
			}
			if str[i] == 'N' {
				x++
			}
			temp := strconv.Itoa(x) + "++++" + strconv.Itoa(y)
			if _, ok := pathDic[temp]; ok {
				flag = false
				continue
			}
			pathDic[temp] = struct{}{}
		}
		if flag {
			fmt.Println("False")
		} else {
			fmt.Println("True")
		}
	}
}
