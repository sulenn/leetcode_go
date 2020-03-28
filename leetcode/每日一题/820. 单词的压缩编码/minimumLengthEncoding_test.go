package main

import (
	"fmt"
	"testing"
)

func Test_minimumLengthEncoding(T *testing.T) {
	fmt.Println(minimumLengthEncoding([]string {"time", "me", "bell"}))
	fmt.Println(minimumLengthEncoding([]string {"my", "myself", "bell"}))
	fmt.Println(minimumLengthEncoding([]string {"time", "me","my", "myself", "bell"}))
}
