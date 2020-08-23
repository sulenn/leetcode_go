package main

import (
	"fmt"
	"io"
	"math/big"
	"strings"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		if n == 0 {
			fmt.Println(0)
			continue
		}
		if n == 1 {
			fmt.Println(1)
			continue
		}
		//if n == 2 {
		//	fmt.Println("3 2")
		//	fmt.Println("1 1")
		//	continue
		//}
		nums := getNums(n)
		result := print(nums, n)
		bigNum := big.NewInt(0)
		for i := 0; i < n; i++ {
			str := ""
			for j := 0; j < n; j++ {
				bigNum = big.NewInt(result[i][j])
				str += bigNum.String() + " "
			}
			str = strings.TrimRight(str, " ")
			fmt.Println(str)
			//str += "\n"
		}

	}
}

func getNums(n int) []int64 {
	var num1 int64 = 1
	var num2 int64 = 1
	result := []int64{1, 1}
	for i := 2; i < n*n; i++ {
		num1, num2 = num2, num2+num1
		result = append(result, num2)
	}
	return result
}

func print(nums []int64, n int) [][]int64 {
	up := 0
	down := n - 1
	left := 0
	right := n - 1
	result := make([][]int64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int64, n)
	}
	index := len(nums) - 1
	for up <= down && left <= right {
		for i := left; i <= right; i++ { // up
			result[up][i] = nums[index]
			index--
		}
		up++
		if down < up {
			break
		}
		for i := up; i <= down; i++ {
			result[i][right] = nums[index]
			index--
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			result[down][i] = nums[index]
			index--
		}
		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			result[i][left] = nums[index]
			index--
		}
		left++
	}
	return result
}
