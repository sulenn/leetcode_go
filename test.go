package main

import "fmt"

func main() {
	test := []int {1,2,3,4,5,6,7}
	for _, v:= range test{
		test = append(test, v)
	}
	fmt.Println(test)
}