package main

import "fmt"

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

func main() {
	//fmt.Println(commonPre([]string {"flower","flow","flight"}))
	//fmt.Println(commonPre([]string {"dog","racecar","car"}))
	//fmt.Println(time.Now().Unix())
	var a string
	fmt.Scanln(&a)
	fmt.Println(a)
}