package main

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{4, 2, 1}))
}
