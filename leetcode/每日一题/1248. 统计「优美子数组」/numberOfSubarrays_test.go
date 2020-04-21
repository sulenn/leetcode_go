package main

import (
	"fmt"
	"testing"
)

func Test_numberOfSubarrays(T *testing.T) {
	fmt.Println(numberOfSubarrays([]int {1,1,2,1,1}, 3))
	fmt.Println(numberOfSubarrays([]int {2,4,6}, 1))
	fmt.Println(numberOfSubarrays([]int {2,2,2,1,2,2,1,2,2,2}, 2))
	fmt.Println(numberOfSubarrays([]int {}, 2))
	fmt.Println(numberOfSubarrays([]int {1,1,1,1,1}, 1))
}
