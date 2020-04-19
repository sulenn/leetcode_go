package main

import (
	"fmt"
	"testing"
)

func Test_jump(T *testing.T) {
	fmt.Println(jump([]int {2,3,1,1,4}))
	fmt.Println(jump([]int {2}))
}
