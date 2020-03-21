package main

import (
	"fmt"
	"testing"
)

func Test_canPartition(T *testing.T) {
	fmt.Println(canPartition([]int {1, 5, 11, 5}))
	fmt.Println(canPartition([]int {1, 2, 3, 5}))
}
