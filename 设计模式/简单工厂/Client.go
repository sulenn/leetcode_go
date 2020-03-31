package main

func main() {
	var simpleFactory = new(SimpleFactory) // 简单工厂类
	var apple Computer = simpleFactory.CreateProduct("Apple")
	var huawei Computer = simpleFactory.CreateProduct("Huawei")
	apple.Brand()
	huawei.Brand()
}
