package main

import (
	"fmt"
	"io"
	"sort"
	"strconv"
)

//https://www.nowcoder.com/practice/de044e89123f4a7482bd2b214a685201?tpId=37&tqId=21231&tPage=1&rp=&ru=%2Fta%2Fhuawei&qru=%2Fta%2Fhuawei%2Fquestion-ranking

func main() {
	for {
		line := 0
		_, err := fmt.Scan(&line) // key，value 行数
		if err == io.EOF {
			break
		}
		dic := make(map[int]int)
		arr := []int{}
		key := 0
		value := 0
		for i := 0; i < line; i++ {
			_, err := fmt.Scan(&key, &value)
			if err == io.EOF {
				break
			}
			if _, ok := dic[key]; ok {
				dic[key] += value
			} else {
				dic[key] = value
				arr = append(arr, key)
			}
		}
		sort.Ints(arr)
		for _, v := range arr {
			fmt.Println(strconv.Itoa(v) + " " + strconv.Itoa(dic[v]))
		}
	}
}
