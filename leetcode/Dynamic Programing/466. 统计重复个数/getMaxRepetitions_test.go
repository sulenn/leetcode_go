package main

import (
	"fmt"
	"testing"
)

func Test_getMaxRepetitions(T *testing.T) {
	fmt.Println(getMaxRepetitions("acb",4,"ab",2))
	fmt.Println(getMaxRepetitions("acbab",4,"ab",2))
}
