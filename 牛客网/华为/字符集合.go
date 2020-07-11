package main

import (
	"fmt"
	"io"
)

//https://www.nowcoder.com/question/next?pid=710802&qid=26013&tid=31505056

func main() {
	for {
		input := ""
		_, err := fmt.Scan(&input)
		if err == io.EOF {
			break
		}
		dic := make(map[byte]struct{})
		arr := []byte{}
		for i := 0; i < len(input); i++ {
			if _, ok := dic[input[i]]; !ok {
				dic[input[i]] = struct{}{}
				arr = append(arr, input[i])
			}
		}
		result := ""
		for _, v := range arr {
			result += string(v)
		}
		fmt.Println(result)
	}
}
