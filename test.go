package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := "123"
	arr := []rune(test)
	fmt.Println(arr)
	var num int = 1
	fmt.Println(arr[0]==1+'0')
	fmt.Println(reflect.TypeOf(arr[0]))
	fmt.Println(reflect.TypeOf(1+'0'))
	fmt.Println(reflect.TypeOf('0'))
	fmt.Println(reflect.TypeOf(test[1]))
	fmt.Println(test[1]=='2')
	fmt.Println(test[1]==50)
	fmt.Println(arr[1]==50)
	fmt.Println('2'==50)
	var num1 int = 50
	//fmt.Println(test[1]==num1)
	//fmt.Println(arr[1]==num1)
	fmt.Println('2'==num1)
	//fmt.Println(arr[0]==num+'0')
	fmt.Println(num+'0')
	fmt.Println(reflect.TypeOf(num+'0'))
	fmt.Println(string(num+'0'))
	fmt.Println(string(49))
}