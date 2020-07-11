package main

import (
	"fmt"
	"testing"
)

func Test_verifyPostorder(T *testing.T) {
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
	fmt.Println(verifyPostorder([]int{1, 3, 2, 5}))
	fmt.Println(verifyPostorder([]int{3, 2, 5, 1}))
}
