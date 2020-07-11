package main

import (
	"fmt"
	"testing"
)

func Test_search(T *testing.T) {
	fmt.Println(search([]int{3, 3, 3}, 3))
	fmt.Println(search([]int{1, 3, 3, 3}, 3))
	fmt.Println(search([]int{1, 3, 3, 3, 1}, 3))
	fmt.Println(search([]int{1}, 3))
	fmt.Println(search([]int{}, 3))
}
