package main

import (
	"fmt"
	"testing"
)

func Test_minNumber(T *testing.T) {
	fmt.Println(minNumber([]int{10, 2}))
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
}
