package main

import (
	"fmt"
	"io"
)

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	var n, m int
	var str string
	for {
		_, err := fmt.Scan(&n, &m, &str)
		if err == io.EOF {
			break
		}
		count := 0 // 满足两个.连在一起的数量，比如“..”为1，“...”为2
		var word byte
		for i := 0; i < n; i++ {
			if word == '.' && str[i] == '.' {
				count++
			} else {
				word = str[i]
			}
		}
		var a int
		var b uint8
		arr := make([]byte, n)
		for i := 0; i < n; i++ {
			arr[i] = str[i]
		}
		for j := 0; j < m; j++ {
			fmt.Scan(&a, &b)
			if b == '.' {
				if arr[a-1] == '.' {
					continue
				}
				arr[a-1] = b
				if a-2 >= 0 && arr[a-2] == '.' {
					count++
				}
				if a < n && arr[a] == '.' {
					count++
				}
			}
			if b != '.' {
				if arr[a-1] == '.' {
					if a-2 >= 0 && arr[a-2] == '.' {
						count--
					}
					if a < n && arr[a] == '.' {
						count--
					}
				}
				arr[a-1] = b
			}
			fmt.Println(count)
		}
	}
}
