package main

import (
	"fmt"
	"testing"
)

func Test_subsetsWithDup(T *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2, 1}))
	fmt.Println(subsetsWithDup([]int{1, 1, 1, 1, 1}))
	fmt.Println(subsetsWithDup([]int{}))
	fmt.Println(subsetsWithDup([]int{2, 1, 2, 1, 3}))
}
