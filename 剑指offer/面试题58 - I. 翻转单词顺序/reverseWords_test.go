package main

import (
	"fmt"
	"testing"
)

func Test_reverseWords(T *testing.T) {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))
}
