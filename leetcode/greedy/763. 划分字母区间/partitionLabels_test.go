package main

import (
	"fmt"
	"testing"
)

func Test_partitionLabels(T *testing.T) {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("aaaaaaaaaa"))
	fmt.Println(partitionLabels("aaaaaaaaaabbbb"))
	fmt.Println(partitionLabels("aaaaaaaaaabbbbaa"))
	fmt.Println(partitionLabels("abcdefg"))
	fmt.Println(partitionLabels(""))
}
