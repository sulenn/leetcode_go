package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		//fmt.Println(str)
		//_, err := fmt.Scan(&input)
		//fmt.Println(input)
		//for err == io.EOF {
		//	break
		//}
		alph := make([]int, 128)
		for i:=0; i<len(str); i++ {
			if int(str[i]) >= 0 && int(str[i])<=127 {
				alph[int(str[i])]++
			}
		}
		result := 0
		for _, v:= range alph {
			if v > 0 {
				result++
			}
		}
		fmt.Println(result)
	}
}
