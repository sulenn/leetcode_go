package main

type SimpleFactory struct {  // 简单工厂类
	
}

func (this *SimpleFactory) CreateProduct(name string) Computer {
	if name == "Apple" {
		return new(Apple)
	} else if name == "Huawei" {
		return new(Huawei)
	} else {
		return new(Other)
	}
}
