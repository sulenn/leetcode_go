package main

import (
	"fmt"
	"testing"
)

func Test_canMeasureWater(T *testing.T) {
	fmt.Println(canMeasureWater(3, 5, 4))
	fmt.Println(canMeasureWater(2, 6, 5))
	fmt.Println(canMeasureWater(1, 0, 1))
}
