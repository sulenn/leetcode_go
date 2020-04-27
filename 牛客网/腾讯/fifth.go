package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var times int
	var a, b int
	for {
		_, err := fmt.Scan(&times)
		if err == io.EOF {break}
		for i:=0; i<times; i++ {
			fmt.Scan(&a, &b)
			//b1 := int(math.Pow(2.0, float64(b)-1))-1
			b2 := int(math.Pow(2.0, float64(b)))-1
			if b == 1 {
				fmt.Println(1)
				continue
			}
			if a < b2 {
				fmt.Println(-1)
				continue
			}
			for a > b2 {
				a /= 2
			}
			fmt.Println(a)
		}
	}
}
