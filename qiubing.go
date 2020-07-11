package main

import (
	"fmt"
	"io"
)

func solution(str string) int {
	chars := []rune(str)
	word := make(map[rune]int)
	for i := 0; i < len(chars); {
		var j int
		for j = i + 1; j < len(chars); j++ {
			if chars[i] != chars[j] {
				break
			}
		}

		ll := j - i
		if v, ok := word[chars[i]]; !ok {
			word[chars[i]] = ll
		} else if ll > v {
			word[chars[i]] = ll
		}

		i = j
	}

	sum := 0
	for _, v := range word {
		sum += v
	}

	return sum
}

func main() {
	var str string
	var res int
	for {
		_, err := fmt.Scanf("%s", &str)
		if err == io.EOF {
			break
		} else {
			res = solution(str)
			fmt.Println(res)
		}
	}
}
