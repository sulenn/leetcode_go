package main

import (
	"fmt"
	"testing"
)

func Test_findLongestWord(T *testing.T) {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "cplea", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "cplea", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
	fmt.Println(findLongestWord("abpcplea", []string{"w"}))
	fmt.Println(findLongestWord("abpcplea", []string{""}))
}
