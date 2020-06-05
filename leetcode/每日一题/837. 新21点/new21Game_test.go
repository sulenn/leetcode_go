package main

import (
	"fmt"
	"testing"
)

func Test_new21Game(T *testing.T) {
	fmt.Println(new21Game(10,1,10))
	fmt.Println(new21Game(6,1,10))
	fmt.Println(new21Game(4,2,10))
	fmt.Println(new21Game(21,17,10))
}
