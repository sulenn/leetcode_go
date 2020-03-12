package main

import (
	"fmt"
	"testing"
)

func Test_gcdOfStrings(T *testing.T) {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABABAB", "ABABA"))
	fmt.Println(gcdOfStrings("ABABAB", "ABABB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
}
