package main

import (
	"fmt"
	"testing"
)

func Test_checkValidString(T *testing.T) {
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*))"))
	fmt.Println(checkValidString("(*)))"))
	fmt.Println(checkValidString("(*))("))
}
