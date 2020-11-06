package main

//https://leetcode-cn.com/problems/backspace-string-compare/

func backspaceCompare(S string, T string) bool {
	sBackNum := 0
	tBackNum := 0
	lengthS := len(S) - 1
	lengthT := len(T) - 1
	for lengthS >= 0 || lengthT >= 0 {
		for lengthS >= 0 {
			if S[lengthS] == '#' {
				lengthS--
				sBackNum++
			} else if sBackNum > 0 {
				lengthS--
				sBackNum--
			} else {
				break
			}
		}
		for lengthT >= 0 {
			if T[lengthT] == '#' {
				lengthT--
				tBackNum++
			} else if tBackNum > 0 {
				lengthT--
				tBackNum--
			} else {
				break
			}
		}
		if lengthS >= 0 && lengthT >= 0 {
			if S[lengthS] != T[lengthT] {
				return false
			}
		} else {
			if lengthS >= 0 || lengthT >= 0 {
				return false
			}
		}
		lengthS--
		lengthT--
	}
	return true
}

func main() {
	println(backspaceCompare("ab##", "c#d#"))
}
