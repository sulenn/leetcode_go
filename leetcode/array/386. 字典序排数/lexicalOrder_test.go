package main

import (
	"fmt"
	"testing"
)

func Test_lexicalOrder(T *testing.T) {
	fmt.Println(lexicalOrder(13))
	fmt.Println(lexicalOrder(113))
	fmt.Println(lexicalOrder(1113))
}
