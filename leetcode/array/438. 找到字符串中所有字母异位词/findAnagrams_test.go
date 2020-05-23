package main

import (
	"fmt"
	"testing"
)

func Test_findAnagrams(T *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
