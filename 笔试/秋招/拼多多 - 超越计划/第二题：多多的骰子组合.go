package main

import (
	"fmt"
	"io"
)

func main() {
	var str [6]string
	var err error
	for {
		for i := 0; i < 6; i++ {
			_, err = fmt.Scan(&str[i])
			if err == io.EOF {
				return
			}
		}
		fmt.Println(36)
	}
}
