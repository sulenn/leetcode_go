package main

import (
	"fmt"
	"testing"
)

func Test_isMatch(T *testing.T) {
	fmt.Println(isMatch("aa","a"))
	fmt.Println(isMatch("aa","a*"))
	fmt.Println(isMatch("aab","c*a*b"))
	fmt.Println(isMatch("mississippi","mis*is*p*."))
	fmt.Println(isMatch("ssippi","s*p*."))
	fmt.Println(isMatch("ab",".*"))
	fmt.Println(isMatch("a","..*"))
}
