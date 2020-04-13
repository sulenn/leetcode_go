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
	obj := Construc([]int{1,2,3,4,5,6,7})
	fmt.Println(obj.arr)
	obj.add(9)
	fmt.Println(obj.arr)
}