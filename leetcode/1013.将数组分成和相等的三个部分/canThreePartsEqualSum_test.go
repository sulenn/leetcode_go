package main

import (
	"fmt"
	"testing"
)

func Test_canThreePartsEqualSum(T *testing.T) {
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
	fmt.Println(canThreePartsEqualSum([]int{1, -1, 1, -1}))
	fmt.Println(canThreePartsEqualSum([]int{10, -10, 10, -10, 10, -10, 10, -10}))
}
