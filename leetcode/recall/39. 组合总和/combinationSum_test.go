package main

import (
	"fmt"
	"testing"
)

func Test_combinationSum(T *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{}, 8))
}
