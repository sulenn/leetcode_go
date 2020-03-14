package main

//https://www.nowcoder.com/questionTerminal/dae9959d6df7466d9a1f6d70d6a11417

import (
"fmt"
"io"
)

func main() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err == io.EOF {
			break
		}
		fmt.Println(a + b)
	}
}

//func main() {
//	a:=0
//	b:=0
//	for {
//		n, _ := fmt.Scan(&a,&b)
//		if n == 0 {
//			break
//		} else {
//			fmt.Printf("%d\n",a+b)
//		}
//	}
//}
