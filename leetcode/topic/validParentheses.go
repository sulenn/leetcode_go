package main

import "fmt"

func isValid(s string) bool {
	stack := make([]uint8,len(s))
	top := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack[top] = s[i]
			top++
		} else if top == 0 {
				return false
		} else if s[i] == ')' && stack[top - 1] == '('{
			top--
		} else if s[i] == '}' && stack[top - 1] == '{' {
				top--
		} else if s[i] == ']' && stack[top - 1] == '[' {
					top --
		} else {
						return false
		}
	}
	return top == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid( "()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid( "([)]"))
	fmt.Println(isValid( "{[]}"))
}

