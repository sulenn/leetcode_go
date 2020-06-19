package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	bytes := []byte(s)
	length := len(bytes)
	pre := 0
	tail := length-1
	for pre < tail {
		preBool := (bytes[pre] >= '0' && bytes[pre] <= '9') || (bytes[pre] >= 'a' && bytes[pre] <= 'z')
		if !preBool {
			pre++
			continue
		}
		tailBool := (bytes[tail] >= '0' && bytes[tail] <= '9') || (bytes[tail] >= 'a' && bytes[tail] <= 'z')
		if !tailBool {
			tail--
			continue
		}
		if bytes[pre] != bytes[tail] {
			return false
		}
		pre++
		tail--
	}
	return true
}
