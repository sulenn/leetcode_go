package main

type Subject interface { // 主题类
	RegisterObserver(observer Observer) // 注册观察者
	RemoveObserver(observer Observer)   // 移除观察者
	NotifyObserver()                    // 通知所有观察者
}

type ConcreteSubject struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
	Observers   []Observer
}

func (this *ConcreteSubject) New() { // 初始化
	this.Observers = make([]Observer, 0)
}

func (this *ConcreteSubject) RegisterObserver(observer Observer) {
	this.Observers = append(this.Observers, observer)
}

func (this *ConcreteSubject) RemoveObserver(observer Observer) {
	for i := 0; i < len(this.Observers); i++ {
		if this.Observers[i] == observer {
			this.Observers = append(this.Observers[:i], this.Observers[i+1:]...)
		}
	}
}

func (this *ConcreteSubject) SetMeasurements(t, h, p float64) { // 更新主题属性
	this.Temperature = t
	this.Humidity = h
	this.Pressure = p
	this.NotifyObserver()
}

func (this *ConcreteSubject) NotifyObserver() {
	for i := 0; i < len(this.Observers); i++ {
		twmp := this.Observers[i]
		twmp.Update(this.Temperature, this.Humidity, this.Pressure)
	}
}
