package main

import (
	"fmt"
	"testing"
)

func Test_eraseOverlapIntervals(T *testing.T) {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
	fmt.Println(eraseOverlapIntervals([][]int{}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}))
}
