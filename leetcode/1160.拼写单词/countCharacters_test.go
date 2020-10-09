package main

import (
	"fmt"
	"testing"
)

func Test_countCharacters(T *testing.T) {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, ""))
	fmt.Println(countCharacters([]string{}, "welldonehoneyr"))
}
