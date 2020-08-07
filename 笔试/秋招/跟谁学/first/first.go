package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

//时间限制： 3000MS
//内存限制： 589824KB
//题目描述：
//给定一个正整数N，寻找[2, N]区间内没有素数的最长子区间，如果最长子区间不止一个，输出任意一个子区间即可
//
//
//
//输入描述
//输入一行，包含一个整数，代表待测区间[2,N]的上界N
//
//输出描述
//输出一行，包含两个整数，代表区间[2,N]内没有素数的最长子区间，子区间为闭区间

func main() {
	var N int
	for {
		_, err := fmt.Scan(&N)
		if err == io.EOF {
			break
		}
		if N < 2 {
			fmt.Println("0 0")
			continue
		}
		if N == 3 {
			fmt.Println("0 0")
			continue
		}
		if N == 3 {
			fmt.Println("0 0")
			continue
		}
		if N == 4 {
			fmt.Println("4 4")
			continue
		}
		if N == 5 {
			fmt.Println("4 4")
			continue
		}
		nums := [2]int{4, 4}
		start := 6
		for i := 6; i <= N; i++ {
			if odd(i) && odd(start) {
				if i-start > nums[1]-nums[0] {
					nums[1] = i
					nums[0] = start
				}
			} else {
				start = i
			}
		}
		fmt.Println(strconv.Itoa(nums[0]) + " " + strconv.Itoa(nums[1]))
	}
}

func odd(num int) bool {
	if num%2 == 0 || num%3 == 0 || num%5 == 0 {
		return true
	}
	value := math.Sqrt(float64(num))
	if int(value)*int(value) == num {
		return true
	}
	return false
}
