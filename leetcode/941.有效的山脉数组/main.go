package main

// https://leetcode-cn.com/problems/valid-mountain-array/

// 参考：https://leetcode-cn.com/problems/valid-mountain-array/solution/tu-jie-941-you-xiao-de-shan-mai-shu-zu-shuang-zhi-/
func validMountainArray(A []int) bool {
	length := len(A)
	p1, p2 := 0, length-1
	flag1, flag2 := false, false
	for p1 != p2 {
		if p1+1 < length && A[p1] < A[p1+1] {
			p1++
			flag1 = true
		}
		if p2-1 >= 0 && A[p2-1] > A[p2] {
			p2--
			flag2 = true
		}
		if !flag1 && !flag2 {
			break
		} else {
			flag1, flag2 = false, false
		}
	}
	if p1 != p2 || p1 == 0 || p1 == length-1 { // error if it keeps rising or failling
		return false
	}
	return true
}
