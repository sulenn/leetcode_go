package main

import (
	"fmt"
	"testing"
)

func Test_reconstructQueue(T *testing.T) {
	fmt.Println(reconstructQueue([][]int {{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}))
	fmt.Println(reconstructQueue([][]int {{1,3},{1,1},{1,2}}))
	fmt.Println(reconstructQueue([][]int {}))
}
