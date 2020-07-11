package main

import (
	"fmt"
	"testing"
)

func Test_findMaxForm(T *testing.T) {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}
