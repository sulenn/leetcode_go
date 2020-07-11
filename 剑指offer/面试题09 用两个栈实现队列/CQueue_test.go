package main

import (
	"fmt"
	"testing"
)

func Test_CQueue(T *testing.T) {
	obj := Constructor()
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(5)
	obj.AppendTail(2)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}
