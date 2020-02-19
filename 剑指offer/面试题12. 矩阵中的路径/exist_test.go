package main

import (
	"fmt"
	"testing"
)

func Test_exist(T *testing.T)  {
	fmt.Println(exist([][]byte {{'a','b','c','e'},{'s','f','c','s'},{'a','d','e','e'}}, "bfce"))
	fmt.Println(exist([][]byte {{'a','b','c','e'},{'s','f','c','s'},{'a','d','e','e'}}, "abfb"))
	fmt.Println(exist([][]byte {{'a', 'a'}}, "aa"))
}
