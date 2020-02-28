package 百度

import (
	"fmt"
	"testing"
)

//func Test_reverse(T *testing.T)  {
//	fmt.Println(reverse([]int {1,2,3,4}))
//	fmt.Println(reverse([]int {1,2,3}))
//	fmt.Println(reverse([]int {1}))
//	fmt.Println(reverse([]int {}))
//}

func Test_fun(T *testing.T)  {
	var test1 *Node
	var test2 *Node
	fmt.Println(fun(test1,test2))
	var test3 *Node
	var test4 = &Node{1,nil}
	fmt.Println(fun(test3,test4))
	var test5 = &Node{1,nil}
	test5.Next= &Node{2,nil}
	var test7 = &Node{1,nil}
	test7.Next = &Node{7,nil}
	print(fun(test5,test7))
}