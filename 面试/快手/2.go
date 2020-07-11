package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//1 22 22 33 22 12 45 44 5
//1 22 54 123

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		sArr := strings.Split(str, " ")
		if len(sArr) < 2 {
			fmt.Println(-1)
			continue
		}
		iArr := make([]int, len(sArr)) // 字符串数组转Int数组
		for k, v := range sArr {
			i, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			iArr[k] = i
		}
		if len(iArr) != len(sArr) { // 异常
			fmt.Println(-1)
			continue
		}
		result := []int{}
		first := math.MinInt64
		second := math.MinInt64
		if iArr[1] < iArr[0] {
			result = append(result, 1)
			first = iArr[0]
			second = iArr[1]
		} else {
			first = iArr[1]
			second = iArr[0]
		}
		for i := 2; i < len(iArr); i++ {
			if iArr[i] < first && iArr[i] >= second {
				result = append(result, i)
				second = iArr[i]
			} else if iArr[i] > first {
				first, second = iArr[i], first
			} else if iArr[i] > second {
				second = iArr[i]
			}
		}
		if len(result) == 0 {
			fmt.Println(-1)
			continue
		}
		resultStr := ""
		for _, v := range result {
			resultStr += strconv.Itoa(v) + " "
		}
		fmt.Println(strings.Trim(resultStr, " "))
	}
}
