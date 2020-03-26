package main

import (
	"fmt"
	"testing"
)

func Test_topKFrequent(T *testing.T) {
	fmt.Println(topKFrequent([]int {1,1,1,2,2,3},2))
	fmt.Println(topKFrequent([]int {1,1,1,2,2},2))
	fmt.Println(topKFrequent([]int {1},1))
	fmt.Println(topKFrequent([]int {},0))
}
