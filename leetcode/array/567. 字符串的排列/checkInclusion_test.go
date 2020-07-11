package main

import (
	"fmt"
	"testing"
)

func Test_checkInclusion(T *testing.T) {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}
