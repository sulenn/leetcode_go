package main

import (
	"fmt"
	"strings"
)

type name struct {
	arr []int
}

func Construc(nums []int) *name {
	return &name{nums}
}

func (this *name) add(num int)  {
	//test := this.arr
	this.arr = append(this.arr, num)
	//test = append(test, num)
}

//func main() {
//	var i *int
//	i=new(int)
//	//*i=10
//	fmt.Println(*i)
//
//}

func main() {
	//fmt.Println(commonPre([]string {"flower","flow","flight"}))
	//fmt.Println(commonPre([]string {"dog","racecar","car"}))
	//fmt.Println(time.Now().Unix())
	//var a string
	//fmt.Scanln(&a)
	//fmt.Println(a)
	var test []string
	fmt.Println(test)
	test = append(test,"2")
	fmt.Println(test)
	var test1 = make([]string, 0)
	fmt.Println(test1)
	strings.Repeat()
}