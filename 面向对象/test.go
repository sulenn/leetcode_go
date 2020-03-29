package main

import "fmt"

type test interface {
	name() string
}

type test1 struct {
}

func (this test1) name() string {
	return "小狗"
}

func (this test1) address() string {
	return "长沙"
}

type test2 struct {}

func (this test2) name() string {
	return "小猫"
}

func main() {
	var temp test = test1{}
	fmt.Println(temp.name())
	//fmt.Println(temp.address())   // 多态不能使用接口本身没有定义的方法
	temp = test2{}
	fmt.Println(temp.name())
}
