package main

import (
	"fmt"
)

const (
	Byte FileSize = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
)

type FileSize int64

func main() {
	fmt.Println(Byte)
	fmt.Println(KiByte)

}
