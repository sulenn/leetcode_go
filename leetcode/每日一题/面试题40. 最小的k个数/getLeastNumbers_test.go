package main

import (
	"fmt"
	"testing"
)

func Test_getLeastNumbers(T *testing.T) {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 2))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 3))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 0))
}
