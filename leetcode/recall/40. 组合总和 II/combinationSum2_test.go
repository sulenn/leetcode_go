package main

import (
	"fmt"
	"testing"
)

func Test_combinationSum2(T *testing.T) {
	fmt.Println(combinationSum2([]int {10,1,2,7,6,1,5},8))
	fmt.Println(combinationSum2([]int {2,5,2,1,2},5))
	fmt.Println(combinationSum2([]int {},5))
}
