package main

import (
	"fmt"
	"testing"
)

func Test_permuteUnique(T *testing.T) {
	fmt.Println(permuteUnique([]int {1,1,2}))
	fmt.Println(permuteUnique([]int {1,1,1}))
	fmt.Println(permuteUnique([]int {}))
	fmt.Println(permuteUnique([]int {2,2,1,1}))
}
