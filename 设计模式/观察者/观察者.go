package main

func main() {
	var concreteSubject = new(ConcreteSubject)
	concreteSubject.New()  // 主题类初始化
	var concreteObserver1 = new(ConcreteObserver1)
	concreteObserver1.RegisterToSubject(concreteSubject)  // 注册观察者1
	var concreteObserver2 = new(ConcreteObserver2)
	concreteObserver2.RegisterToSubject(concreteSubject)  // 注册观察者2
	concreteSubject.SetMeasurements(0,0,0)  // 主题更新，通知观察者
	concreteSubject.SetMeasurements(1,1,1)  // 主题更新，通知观察者
}
