package main

import (
	"fmt"
	"testing"
)

func Test_frequencySort(T *testing.T) {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort(""))
	fmt.Println(frequencySort("aaaaaaaaaaaaaaa"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}
