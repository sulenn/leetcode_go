package main

import (
	"fmt"
	"testing"
)

func Test_myPow(T *testing.T) {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2.0, -2))
	fmt.Println(myPow(0.0, -2))
	fmt.Println(myPow(0.0, 2))
	fmt.Println(myPow(0.0, 0))
	fmt.Println(myPow(0.00000001, -2))
}
