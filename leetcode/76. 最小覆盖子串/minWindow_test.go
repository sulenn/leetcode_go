package main

import (
	"fmt"
	"testing"
)

func Test_minWindow(T *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("aa", "aa"))
}
