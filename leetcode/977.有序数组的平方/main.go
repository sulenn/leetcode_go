package main

//https://leetcode-cn.com/problems/squares-of-a-sorted-array/

func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	prev := 0
	tail := len(A) - 1
	index := len(A) - 1
	for prev <= tail {
		if A[prev]*A[prev] > A[tail]*A[tail] {
			result[index] = A[prev] * A[prev]
			prev++
		} else {
			result[index] = A[tail] * A[tail]
			tail--
		}
		index--
	}
	return result
}
