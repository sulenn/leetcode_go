package main

import (
	"fmt"
	"testing"
)

func Test_fourSum(T *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2},0))
	fmt.Println(fourSum([]int{0,0,0,0,0,0,0,0},0))
	fmt.Println(fourSum([]int{-1,-1,-1,-1,0,0,0,0,0,0,0,0,1,1,1,1},0))
	fmt.Println(fourSum([]int{-3,-1,0,2,4,5},2))
	fmt.Println(fourSum([]int{-1,0,1,2,-1,-4},-1))
}
