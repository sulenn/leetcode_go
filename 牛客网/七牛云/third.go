package main

import "fmt"

/**
 * 将常规记法算术式转为逆波兰表达式形式，输入为一个字符串数组，每一个元素为一个数字或者符号，将每个元素顺序排列即得到待转换的算术式。输出为一个字符串数组，将每个元素顺序排列即得到转换后的逆波兰表达式。
 * @param input string字符串一维数组 待转换的算术式
 * @return string字符串一维数组
 */
func ReversePolishNotation2( input []string ) []string {
	// write code here
	result := []string {}
	stack := []string {}
	for i:=0; i<len(input);i++ {
		if input[i] == ")" {
			stack = stack[:len(stack)-1]
			j:=len(stack)-1
			for ; j>=0 && stack[j] != "("; j-- {
				result = append(result, stack[j])
			}
			stack = stack[:j+1]
		} else if input[i] == "(" || input[i] == "+" || input[i] == "-" || input[i] == "*" || input[i] == "/" {
			stack = append(stack, input[i])
		} else {
			result = append(result, input[i])
			if i+1 < len(input) {
				if input[i+1] == "*" || input[i+1] == "/" {
					if len(stack) > 0 && stack[len(stack)-1] != "*" && stack[len(stack)-1] != "/" && stack[len(stack)-1] != "(" {
						continue
					}
				}
				j:=len(stack)-1
				for ; j>=0&&stack[j] != "("; j-- {
					result = append(result, stack[j])
				}
				stack = stack[:j+1]

			}
		}
	}
	for j:=len(stack)-1; j>=0; j-- {
		result = append(result, stack[j])
	}
	return result
}
//["1","+","2"]
//["3", "-", "(", "4", "*", "5", "+", "1", ")"]
func main() {
	fmt.Println(ReversePolishNotation2([]string{"1","+","2"}))
	fmt.Println(ReversePolishNotation2([]string{"1","+","2","*","6","-","2"}))
	fmt.Println(ReversePolishNotation2([]string{"1","*","2","*","4"}))
	fmt.Println(ReversePolishNotation2([]string{"1","*","2","*","4","*","7"}))
	fmt.Println(ReversePolishNotation2([]string{"1","*","2","*","(","4","*","7",")"}))
	fmt.Println(ReversePolishNotation2([]string{"1","/","2","*","4"}))
	fmt.Println(ReversePolishNotation2([]string{"1","*","2","/","4"}))
	fmt.Println(ReversePolishNotation2([]string{"1"}))
	fmt.Println(ReversePolishNotation2([]string{}))
	fmt.Println(ReversePolishNotation2([]string{"1","*","2","*","3"}))
	fmt.Println(ReversePolishNotation2([]string{"1","*","2","+","5","/","8"}))
	fmt.Println(ReversePolishNotation2([]string{"1","*","(","2","+","5",")","/","8"}))
	fmt.Println(ReversePolishNotation2([]string{"3","-","(","4","*","5","+","1",")"}))
	fmt.Println(ReversePolishNotation2([]string{"3","-","(","4","*","(","5","+","1",")",")"}))
	fmt.Println(ReversePolishNotation2([]string{"3","-","(","4","*","(","(","5","+","1",")",")",")"}))
	fmt.Println(ReversePolishNotation2([]string{"(","3","-","(","4","*","(","(","5","+","1",")",")",")",")"}))
}