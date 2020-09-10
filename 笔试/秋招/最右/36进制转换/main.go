package main

import (
	"fmt"
	"io"
)

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		fmt.Println(tranfer(str))
	}

}

func tranfer(str string) int64 {
	if len(str) == 0 {
		return 0
	}
	falg := false
	if str[0] == '-' {
		falg = true
		str = str[1:]
	}
	if len(str) == 0 {
		return 0
	}
	var base int64 = 1
	var num int64 = 0
	if str[len(str)-1] < '0' || (str[len(str)-1] > '9' && str[len(str)-1] < 'a') || str[len(str)-1] > 'z' {
		return 0
	}
	if str[len(str)-1] <= '9' {
		num = int64(str[len(str)-1] - '0')
	} else {
		num = int64(str[len(str)-1] - 'a' + 10)
	}
	var result = num
	for i := len(str) - 2; i >= 0; i-- {
		base *= 36
		if str[i] < '0' || (str[i] > '9' && str[i] < 'a') || str[0] > 'z' {
			return 0
		}
		if base < 0 {
			if falg {
				return 9223372036854775807
			} else {
				return -9223372036854775807
			}
		}

		if str[i] <= '9' {
			num = int64(str[i] - '0')
		} else {
			num = int64(str[i] - 'a' + 10)
		}
		num *= base
		result += num
	}
	if falg {
		return -result
	} else {
		return result
	}
}

//func add(str1 string, str2 string) string {
//	str := ""
//	p1 := len(str1) - 1
//	p2 := len(str2) - 1
//	rest := 0
//	for p1 >= 0 && p2 >= 0 {
//		sum := int(str1[p1]-'0'+str2[p2]-'0') + rest
//		rest = sum / 10
//		str += strconv.Itoa(sum % 10)
//		p1--
//		p2--
//	}
//	for p1 >= 0 {
//		sum := int(str1[p1]-'0') + rest
//		rest = sum / 10
//		str += strconv.Itoa(sum % 10)
//		p1--
//	}
//	for p2 >= 0 {
//		sum := int(str2[p2]-'0') + rest
//		rest = sum / 10
//		str += strconv.Itoa(sum % 10)
//		p2--
//	}
//	return str
//}
