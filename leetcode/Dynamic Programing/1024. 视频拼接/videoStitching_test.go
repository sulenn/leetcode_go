package main

import (
	"fmt"
	"testing"
)

func Test_videoStitching(T *testing.T) {
	fmt.Println(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10))
	fmt.Println(videoStitching([][]int{{0, 1}, {1, 2}}, 5))
}
