package main

//https://leetcode-cn.com/problems/longest-mountain-in-array/

func longestMountain(A []int) int {
	start := -1
	max := 0
	flag := 2 //0 is up, 1 is down, 2 is equal or start point
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			flag = 2
		}
		if A[i] > A[i-1] && flag != 0 {
			flag = 0
			start = i - 1
		}
		if A[i] < A[i-1] && flag != 2 {
			if i-start+1 > max {
				max = i - start + 1
			}
			flag = 1
		}
	}
	return max
}
