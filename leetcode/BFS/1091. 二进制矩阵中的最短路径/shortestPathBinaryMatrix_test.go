package main

import (
	"fmt"
	"testing"
)

func Test_shortestPathBinaryMatrix(T *testing.T) {
	fmt.Println(shortestPathBinaryMatrix([][]int {{0,1},{1,0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int {{0,0,0},{1,1,0},{1,1,0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int {{0}}))
}
