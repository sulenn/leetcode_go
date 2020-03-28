package main

import (
	"fmt"
	"io"
)

func main() {
	var str string
	var res int
	for {
		_, err := fmt.Scanf("%s", &str)
		if err == io.EOF {
			break
		} else {
			res = way(str)
			fmt.Println(res)
		}
	}
}

func way(str string) int {
	dic := make(map[string]struct{})
	cur := 0
	for i:=0;i <len(str);i++ {
		if str[cur] != str[i] {
			cur = i
		}
		if _,ok:= dic[str[cur:i+1]]; !ok {
			dic[str[cur:i+1]] = struct{}{}
		}
	}
	if _,ok:= dic[str[cur:]]; !ok {
		dic[str[cur:]] = struct{}{}
	}
	return len(dic)
}
