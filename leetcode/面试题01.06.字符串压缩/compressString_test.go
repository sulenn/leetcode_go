package main

import (
	"fmt"
	"testing"
)

func Test_compressString(T *testing.T) {
	fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(compressString("abbccd"))
}
