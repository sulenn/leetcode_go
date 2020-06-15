package main

import (
	"fmt"
	"testing"
)

func Test_dailyTemperatures(T *testing.T)  {
	fmt.Println(dailyTemperatures([]int {55,38,53,81,61,93,97,32,43,78}))
	//fmt.Println(dailyTemperatures([]int {73, 74, 75, 71, 69, 72, 76, 73}))
}
