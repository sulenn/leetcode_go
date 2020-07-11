package main

import "fmt"

type Computer interface { // 产品类，下面是产品类的三个子类
	Brand()
}

type Apple struct {
	Name string
}

func (this *Apple) Brand() {
	this.Name = "Apple"
	fmt.Println("This Computer Brand is ", this.Name)
}

type Huawei struct {
	Name string
}

func (this *Huawei) Brand() {
	this.Name = "Huawei"
	fmt.Println("This Computer Brand is ", this.Name)
}

type Other struct {
	Name string
}

func (this *Other) Brand() {
	this.Name = "Nothing"
	fmt.Println("This Computer Brand is ", this.Name)
}
