package main

import (
	"fmt"
	"testing"
)

func Test_maximalSquare(T *testing.T) {
	fmt.Println(maximalSquare([][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}))
}
