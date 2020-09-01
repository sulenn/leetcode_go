package main

import (
	"fmt"
	"testing"
)

func Test_quickSort(T *testing.T) {
	test1 := []int{2, 1, 5, 4, 7, 6, 8, 0}
	sort(test1)
	fmt.Println(test1)
	test2 := []int{2, 1, 6, 3, 76, 8, 444, 87, 93, 0}
	quickSort(test2)
	fmt.Println(test2)
}
