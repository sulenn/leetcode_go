package main

import (
	"fmt"
	"testing"
)

func Test_merge(T *testing.T) {
	test1A := []int{1, 2, 3, 0, 0, 0}
	test1B := []int{2, 5, 6}
	merge(test1A, 3, test1B, 3)
	fmt.Println(test1A)
	test2A := []int{2, 5, 6, 0, 0, 0}
	test2B := []int{1, 2, 3}
	merge(test2A, 3, test2B, 3)
	fmt.Println(test2A)
}
