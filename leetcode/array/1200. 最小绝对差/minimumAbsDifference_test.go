package main

import (
	"fmt"
	"testing"
)

func Test_minimumAbsDifference(T *testing.T) {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}
