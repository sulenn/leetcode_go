package main

import (
	"fmt"
	"testing"
)

func Test_cuttingRope(T *testing.T) {
	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(3))
	fmt.Println(cuttingRope(10))
	fmt.Println(cuttingRope(-1))
}
