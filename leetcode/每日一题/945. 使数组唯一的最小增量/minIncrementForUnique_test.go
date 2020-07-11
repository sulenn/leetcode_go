package main

import (
	"fmt"
	"testing"
)

func Test_minIncrementForUnique(T *testing.T) {
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}
