package main

import (
	"fmt"
	"testing"
)

func Test_findPaths(T *testing.T) {
	fmt.Println(findPaths(2, 2, 2, 0, 0))
	fmt.Println(findPaths(1, 3, 3, 0, 1))
	fmt.Println(findPaths(8, 50, 23, 5, 26))
	fmt.Println(findPaths(36, 5, 50, 15, 3))
}
