package main

//https://leetcode-cn.com/problems/sorted-merge-lcci/

// O(n) 解法，从 A 数组从后往前遍历
func merge(A []int, m int, B []int, n int) {
	m -= 1
	n -= 1
	for i := m + n + 1; i >= 0; i-- {
		if m < 0 || n < 0 {
			break
		}
		if A[m] > B[n] {
			A[i] = A[m]
			m--
		} else {
			A[i] = B[n]
			n--
		}
	}
	if n >= 0 { // B 数组中还有数字未被遍历
		A = append(A[:0], B[:n+1]...)
	}
}
