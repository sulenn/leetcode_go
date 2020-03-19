package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.nowcoder.com/practice/119bcca3befb405fbe58abe9c532eb29?tpId=37&tqId=21240&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		arr := strings.Split(str, ";")
		location := []int {0,0}
		for _, v := range arr {
			if judge1(v) {
				num, _ := strconv.Atoi(v[1:])   // 只考虑了正数
				if v[0] == 'A' {
					location[0] -= num
				}
				if v[0] == 'D' {
					location[0] += num
				}
				if v[0] == 'W' {
					location[1] += num
				}
				if v[0] == 'S' {
					location[1] -= num
				}
			}
		}
		fmt.Println(strconv.Itoa(location[0])+","+strconv.Itoa(location[1]))
	}
}

func judge1(s string) bool {  // 判断每一个操作是否合法，只考虑了正数
	if len(s) < 2 || len(s) >3 {return false}
	if s[0] == 'A' || s[0] == 'D' || s[0] == 'W' || s[0] == 'S' {
		if s[1] >= '0' && s[1] <='9' && len(s) == 2 {
			return true
		} else if s[1] >= '0' && s[1] <='9' && s[2] >= '0' && s[2] <='9' {
			return true
		}
	}
	return false
}
