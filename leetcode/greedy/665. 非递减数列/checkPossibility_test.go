package main

import (
	"fmt"
	"testing"
)

func Test_checkPossibilityT (t *testing.T) {
	fmt.Println(checkPossibility([]int {3,4,2,3}))
	fmt.Println(checkPossibility([]int {4,3,2}))
	fmt.Println(checkPossibility([]int {4,2,3}))
	fmt.Println(checkPossibility([]int {-1,4,2,3}))
	fmt.Println(checkPossibility([]int {4,2,1}))
	fmt.Println(checkPossibility([]int {3,2,2}))
	fmt.Println(checkPossibility([]int {3,3,2}))
}
