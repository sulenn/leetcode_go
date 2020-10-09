package main

import (
	"fmt"
	"github.com/sulenn/leetcode_go/leetcode/每日一题"
	"testing"
)

func Test_minWindow(T *testing.T) {
	fmt.Println(每日一题.minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(每日一题.minWindow("aa", "aa"))
}
