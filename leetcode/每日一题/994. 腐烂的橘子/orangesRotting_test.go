package main

import (
	"fmt"
	"testing"
)

func Test_orangesRotting(T *testing.T)  {
	fmt.Println(orangesRotting([][]int {{2,1,1}, {1,1,0}, {0,1,1}}))
	fmt.Println(orangesRotting([][]int {{2,1,1}, {0,1,1}, {1,0,1}}))
	fmt.Println(orangesRotting([][]int {{0,2}}))
	fmt.Println(orangesRotting([][]int {{0}}))
}