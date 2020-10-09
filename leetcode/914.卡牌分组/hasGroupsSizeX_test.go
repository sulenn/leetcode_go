package main

import (
	"fmt"
	"testing"
)

func Test_hasGroupsSizeX(T *testing.T) {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(hasGroupsSizeX([]int{1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
	fmt.Println(hasGroupsSizeX([]int{}))
}
