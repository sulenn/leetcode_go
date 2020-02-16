package main

import (
	"fmt"
	"testing"
)

func Test_validateStackSequences(T *testing.T)  {
	fmt.Println(validateStackSequences([]int {1,2,3,4,5}, []int {4,5,3,2,1}))
	fmt.Println(validateStackSequences([]int {1,2,3,4,5}, []int {4,3,5,1,2}))
}
