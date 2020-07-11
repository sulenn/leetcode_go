package main

import (
	"fmt"
	"testing"
)

func Test_divide(T *testing.T) {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(-2147483648, -1))
}
