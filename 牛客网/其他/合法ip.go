package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	var ipv4 string
	for {
		_, err := fmt.Scan(&ipv4)
		if err == io.EOF {
			break
		}
		if len(ipv4) > 15 {
			fmt.Println("NO")
			continue
		}
		strs := strings.Split(ipv4, ".")
		flag := false
		for i := 0; i < len(strs); i++ {
			num, err := strconv.Atoi(strs[i])
			if err != nil {
				fmt.Println("NO")
				flag = true
				break
			}
			if num < 0 || num > 255 {
				fmt.Println("NO")
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		fmt.Println("YES")
	}
}
