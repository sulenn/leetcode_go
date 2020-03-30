package main

import "fmt"

//参考：https://blog.csdn.net/cloudUncle/article/details/84865967

type Handler interface {
	SetNext(handler Handler)  // 设置下一个实例
	HandleRequest(t int)  // 处理请求
}

type ConcreteHandler1 struct {   // 实现类1
	Next Handler
}

func (this *ConcreteHandler1) HandleRequest(t int) {
	if t == 1 {
		fmt.Println("request",t," is handled by ConcreteHandler1")
		return
	}
	if this.Next != nil{
		this.Next.HandleRequest(t)
	}
}

func (this *ConcreteHandler1) SetNext(handler Handler)  {
	this.Next = handler
}

type ConcreteHandler2 struct {  // 实现类2
	Next Handler
}

func (this *ConcreteHandler2) HandleRequest(t int) {
	if t == 2 {
		fmt.Println("request",t," is handled by ConcreteHandler1")
		return
	}
	if this.Next != nil{
		this.Next.HandleRequest(t)
	}
}

func (this *ConcreteHandler2) SetNext(handler Handler)  {
	this.Next = handler
}

func main() {  // 测试
	var handler1 Handler = &ConcreteHandler1{nil}
	var handler2 Handler = &ConcreteHandler2{nil}
	handler2.SetNext(handler1)
	handler2.HandleRequest(1)
	handler2.HandleRequest(2)
}