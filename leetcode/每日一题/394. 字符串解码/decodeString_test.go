package main

import (
	"fmt"
	"testing"
)

func Test_decodeString(T *testing.T) {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("ef2[abc]3[cd]ef"))
}
