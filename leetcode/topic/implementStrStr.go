package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	for head := 0; head+len(needle) <= len(haystack); head++ {
		if haystack[head:head+len(needle)] == needle {
			return head
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}
