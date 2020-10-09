package main

import (
	"fmt"
	"testing"
)

func Test_backTracking(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{}))
}
