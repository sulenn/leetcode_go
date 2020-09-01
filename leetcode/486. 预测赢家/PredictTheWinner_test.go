package main

import (
	"fmt"
	"testing"
)

func Test_PredictTheWinner(t *testing.T) {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
}
