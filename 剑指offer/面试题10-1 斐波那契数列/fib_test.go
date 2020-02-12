package main

import (
	"fmt"
	"testing"
)

func Test_fib(T *testing.T)  {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(45))
	fmt.Println(fib(95))
}
