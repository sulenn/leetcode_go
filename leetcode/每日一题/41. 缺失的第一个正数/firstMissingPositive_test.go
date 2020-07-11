package main

import (
	"fmt"
	"testing"
)

func Test_firstMissingPositive(T *testing.T) {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
