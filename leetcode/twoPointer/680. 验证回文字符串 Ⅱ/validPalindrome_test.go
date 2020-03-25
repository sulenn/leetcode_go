package main

import (
	"fmt"
	"testing"
)

func Test_validPalindrome(T *testing.T) {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abcca"))
	fmt.Println(validPalindrome("abbca"))
	fmt.Println(validPalindrome("abbcca"))
	fmt.Println(validPalindrome(""))
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}
