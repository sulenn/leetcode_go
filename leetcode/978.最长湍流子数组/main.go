package main

import "fmt"

//https://leetcode-cn.com/problems/longest-turbulent-subarray/

func maxTurbulenceSize(A []int) int {
	if len(A) == 0 {
		return 0
	}
	flag := 0 // 0 表示相等、1表示大于、2表示小于
	count := 0
	max := 1
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			if flag == 2 {
				count++
			} else {
				count = 2
			}
			flag = 1
		} else if A[i] < A[i-1] {
			if flag == 1 {
				count++
			} else {
				count = 2
			}
			flag = 2
		} else {
			count = 0
			flag = 0
		}
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
	fmt.Println(maxTurbulenceSize([]int{4, 8, 12, 16}))
	fmt.Println(maxTurbulenceSize([]int{100}))
}
