package main

import (
	"fmt"
	"testing"
)

func Test_MaxQueue(T *testing.T) {
	obj := Constructor()
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Max_value())
	obj.Push_back(1)
	obj.Push_back(2)
	fmt.Println(obj.Max_value())
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Max_value())
}
