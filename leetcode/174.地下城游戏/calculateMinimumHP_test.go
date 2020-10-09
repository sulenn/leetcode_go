package main

import (
	"fmt"
	"testing"
)

func Test_calculateMinimumHP(t *testing.T) {
	fmt.Println(calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	fmt.Println(calculateMinimumHP([][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}))
	fmt.Println(calculateMinimumHP([][]int{{0, 0}}))

}
