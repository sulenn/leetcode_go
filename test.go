package main

import "log"

const (
	Byte FileSize = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
)

type FileSize int64

func main() {
	log.Fatal("sdgfsdgsgd")

}
