package main

import (
	"fmt"
	"testing"
)

func Test_findDuplicate(T *testing.T) {
	fmt.Println(findDuplicate([]int {1,3,4,2,2}))
	fmt.Println(findDuplicate([]int {3,1,3,4,2}))
}
