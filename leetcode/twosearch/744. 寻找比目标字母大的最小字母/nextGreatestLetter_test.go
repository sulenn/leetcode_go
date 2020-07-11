package main

import (
	"fmt"
	"testing"
)

func Test_nextGreatestLetter(T *testing.T) {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	fmt.Println(nextGreatestLetter([]byte{'j', 'c', 'f'}, 'a'))
	fmt.Println(nextGreatestLetter([]byte{'j', 'c', 'f'}, 'c'))
	fmt.Println(nextGreatestLetter([]byte{'j', 'c', 'f'}, 'd'))
	fmt.Println(nextGreatestLetter([]byte{'j', 'c', 'f'}, 'g'))
	fmt.Println(nextGreatestLetter([]byte{'j', 'c', 'f'}, 'j'))
	fmt.Println(nextGreatestLetter([]byte{'j', 'c', 'f'}, 'k'))
}
