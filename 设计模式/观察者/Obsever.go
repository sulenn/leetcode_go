package main

import "fmt"

type Observer interface { // 接口是指针类型
	Update(t, h, p float64) // 更新观察者
}

type ConcreteObserver1 struct { // 结构体是值类型，在函数中非指针传递时，属于拷贝，任何修改都不会保留
}

func (this *ConcreteObserver1) Update(t, h, p float64) { // 更新
	fmt.Println("ConcreteObsever1.update: ", t, " ", h, " ", p)
}

func (this *ConcreteObserver1) RegisterToSubject(subject Subject) {
	subject.RegisterObserver(this)
}

type ConcreteObserver2 struct {
}

func (this *ConcreteObserver2) Update(t, h, p float64) {
	fmt.Println("ConcreteObsever2.update: ", t, " ", h, " ", p)
}

func (this *ConcreteObserver2) RegisterToSubject(subject Subject) {
	subject.RegisterObserver(this)
}
