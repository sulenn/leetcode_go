package main

import "fmt"

/**
 * Complement
 * @param A int整型一维数组 第一个数组
 * @param B int整型一维数组 第二个数组
 * @return int整型一维数组
 */
func Complement( A []int ,  B []int ) []int {
	// write code here
	lenA := len(A)
	lenB := len(B)
	i:=0
	j:=0
	for ; i<lenA && j<lenB; {
		if A[i] == B[j] {
			A = append(A[:i], A[i+1:]...)
			lenA--
			j++
		} else if A[i] > B[j] {
				j++
		} else if A[i] < B[j] {
					i++
		}
	}
	return A
}
//[1,2,3,4,5],[2,3]
func main() {
	fmt.Println(Complement([]int {1,2,3,4,5}, []int{2,3}))
	fmt.Println(Complement([]int {1,2,3,4,5}, []int{5,6,7}))
}