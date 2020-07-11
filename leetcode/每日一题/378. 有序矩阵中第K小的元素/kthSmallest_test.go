package main

import (
	"fmt"
	"testing"
)

func Test_kthSmallest(T *testing.T) {
	fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
}
