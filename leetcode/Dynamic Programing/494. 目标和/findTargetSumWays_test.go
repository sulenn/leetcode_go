package main

import (
	"fmt"
	"testing"
)

func Test_findTargetSumWays(T *testing.T) {
	fmt.Println(findTargetSumWays([]int {1,1,1,1,1}, 3))
	fmt.Println(findTargetSumWays([]int {1}, 2))
}
