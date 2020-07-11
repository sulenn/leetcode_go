package main

import (
	"testing"
)

func Test_duplicateZeros(T *testing.T) {
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
	duplicateZeros([]int{1, 0, 2, 3, 0, 0})
	duplicateZeros([]int{1, 2, 3})
}
