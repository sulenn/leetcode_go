package main

import (
	"fmt"
	"testing"
)

func Test_stoneGame(T *testing.T) {
	fmt.Println(stoneGame([]int{5,3,4,5}))
	fmt.Println(stoneGame([]int{5,11,5}))
	fmt.Println(stoneGame([]int{5,9,5}))
}
