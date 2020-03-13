package main

import (
	"fmt"
	"testing"
)

func Test_majorityElement(T *testing.T) {
	fmt.Println(majorityElement([]int {3,2,3}))
	fmt.Println(majorityElement([]int {2,2,1,1,1,2,2}))
}
