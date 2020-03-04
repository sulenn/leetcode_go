package main

import (
	"fmt"
	"testing"
)

func Test_sort(T *testing.T) {
	test1 := []int {2,1}
	sort(test1 ,0 ,1)
	fmt.Println(test1)
	test2 := []int {2,1,0}
	sort(test2,0 ,2)
	fmt.Println(test2)
}
