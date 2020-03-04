package main

import (
	"fmt"
	"testing"
)

func Test_quickSort(T *testing.T)  {
	test1 := []int {2,1,5,4,7,6,8,0}
	quickSort(test1)
	fmt.Println(test1)
}