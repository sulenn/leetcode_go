package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//col,_ := strconv.Atoi(input.Text())
		arr := make([][]int, 3) // 数组
		for i := 0; i < 3; i++ {
			input.Scan()
			tempArr := strings.Split(input.Text(), " ")
			numArr := []int{}
			for _, v := range tempArr {
				num, _ := strconv.Atoi(v)
				numArr = append(numArr, num)
			}
			arr[i] = numArr
		}
		result := math.MaxInt64
		numsArr := [][]int{}
		for i := 0; i < len(arr[0]); i++ {
			if len(numsArr) == 0 {
				numsArr = append(numsArr, []int{arr[0][i]})
				numsArr = append(numsArr, []int{arr[1][i]})
				numsArr = append(numsArr, []int{arr[2][i]})
			} else {
				length := len(numsArr)

				for j := 0; j < length; j++ {
					cur := numsArr[0]
					numsArr = numsArr[1:]
					numsArr = append(numsArr, append(cur, arr[0][i]))
					numsArr = append(numsArr, append(cur, arr[1][i]))
					numsArr = append(numsArr, append(cur, arr[2][i]))
				}
			}
		}
		for i := 0; i < len(numsArr); i++ {
			num := abs(numsArr[i])
			if num < result {
				result = num
			}
		}
		fmt.Println(result)
	}
}

func abs(arr []int) int {
	result := 0
	for i := 0; i < len(arr)-1; i++ {
		result += minusAbs(arr[i+1], arr[i])
	}
	return result
}

func minusAbs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
