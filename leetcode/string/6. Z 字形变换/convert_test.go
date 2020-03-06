package main

import (
	"fmt"
	"testing"
)

func Test_convert(T *testing.T) {
	fmt.Println(convert("LEETCODEISHIRING", 3))
	fmt.Println(convert("LEETCODEISHIRING", 4))
	fmt.Println(convert("AB", 1))
}
