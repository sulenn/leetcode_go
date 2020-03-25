package main

import (
	"fmt"
	"testing"
)

func Test_surfaceArea(T *testing.T) {
	fmt.Println(surfaceArea([][]int {{2}}))
	fmt.Println(surfaceArea([][]int {{1,2},{3,4}}))
	fmt.Println(surfaceArea([][]int {{1,0},{0,1}}))
	fmt.Println(surfaceArea([][]int {{1,1,1},{1,0,1},{1,1,1}}))
	fmt.Println(surfaceArea([][]int {{2,2,2},{2,1,2},{2,2,2}}))
}
