package main

import (
	"fmt"
	"testing"
)

func Test_sortColors(T *testing.T) {
	temp := []int {2,0,2,1,1,0}
	sortColors(temp)
	fmt.Println(temp)
	temp = []int {0,0,0,0}
	sortColors(temp)
	fmt.Println(temp)
	temp = []int {1,1,1,1}
	sortColors(temp)
	fmt.Println(temp)
	temp = []int {2,2,2,2}
	sortColors(temp)
	fmt.Println(temp)
	temp = []int {}
	sortColors(temp)
	fmt.Println(temp)
}
