package main

import (
	"fmt"
	"io"
)

func main() {
	var num int
	//_, err := fmt.Scanf("%d", &num)
	_, err := fmt.Scanln(&num)
	if err == io.EOF {
		return
	}
	var count int
	for i:=0; i<num; i++ {
		_, err := fmt.Scanf("%d", &count)
		//_, err := fmt.Scanln(&count)
		if err == io.EOF {
			break
		}
		p1 := 2.0
		p2 := 1.0
		var sum float64
		for j:=0; j<count;j++ {
			sum += p1/p2
			p1, p2 = p1+p2, p1
		}
		fmt.Printf("%.4f",sum)
	}
}
