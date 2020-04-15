package main

import (
	"fmt"
	"testing"
)

func Test_updateMatrix(T *testing.T) {
	fmt.Println(updateMatrix([][]int {{0,0,0},{0,1,0},{0,0,0}}))
	fmt.Println(updateMatrix([][]int {{0,0,0},{0,1,0},{1,1,1}}))
	fmt.Println(updateMatrix([][]int {{0,1,1},{1,1,1},{1,1,1}}))
}
