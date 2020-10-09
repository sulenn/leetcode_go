package main

import (
	"fmt"
	"testing"
)

func Test_canJump(T *testing.T) {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 5, 4}))
	fmt.Println(canJump([]int{}))
	fmt.Println(canJump([]int{2, 5, 0, 0}))
}
