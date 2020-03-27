package main

import (
	"fmt"
	"testing"
)

func Test_searchRange(T *testing.T) {
	fmt.Println(searchRange([]int{8,8}, 8))
	fmt.Println(searchRange([]int{8,8,8}, 8))
	fmt.Println(searchRange([]int{7,9}, 8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
}
