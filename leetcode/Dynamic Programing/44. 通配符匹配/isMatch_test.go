package main

import (
	"fmt"
	"testing"
)

func Test_isMatch(T *testing.T) {
	//fmt.Println(isMatch("aa", "a"))
	//fmt.Println(isMatch("aa", "*"))
	//fmt.Println(isMatch("", "*"))
	//fmt.Println(isMatch("cb", "?a"))
	//fmt.Println(isMatch("adceb", "*a*b"))
	//fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("zacabz", "*a?b*"))
	fmt.Println(isMatch("", ""))
}
