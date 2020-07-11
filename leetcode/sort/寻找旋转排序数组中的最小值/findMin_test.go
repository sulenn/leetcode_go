package main

import (
	"fmt"
	"testing"
)

func Test_findMin(T *testing.T) {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2, 3}))
	fmt.Println(findMin([]int{4, 5, 6, 0, 1, 2, 3}))
	fmt.Println(findMin([]int{4, 0, 1, 2, 3}))
	fmt.Println(findMin([]int{0, 1, 2, 3, -1}))
	fmt.Println(findMin([]int{0, 1, 2, 3}))
}
