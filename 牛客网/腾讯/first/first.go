package main

import (
	"fmt"
	"io"
)

func main() {
	var num int
	for {
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}
			arr:= make([]int, 0)
			line := ""
			for j:=0; j<num; j++ {
				fmt.Scan(&line)
				if line == "peek" {
					if len(arr) != 0 {
						fmt.Println(arr[0])
					} else {
						fmt.Println(-1)
					}
				} else if line == "poll" {
					if len(arr) != 0 {
						//fmt.Println(arr[len(arr)-1])
						arr = arr[1:]
					} else {
						fmt.Println(-1)
					}
				} else if line == "add" {
					var number int
					fmt.Scan(&number)
					arr = append(arr, number)
				}
			}
		}
}

//1
//7
//PUSH 1
//PUSH 2
//TOP
//POP
//TOP
//POP
//POP
//5
//PUSH 1
//PUSH 2
//SIZE
//POP
//SIZE