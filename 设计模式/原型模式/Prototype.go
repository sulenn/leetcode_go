package main

type Prototype interface {
	Clone() Prototype
	Construct(s string)
	GetString() string
}

type ConcreteProtype struct {
	s string
}

func (this *ConcreteProtype) Construct(s string) {
	this.s = s
}

func (this *ConcreteProtype) Clone() Prototype {
	var clone = new(ConcreteProtype)
	clone.Construct(this.s)
	return clone
}

func (this *ConcreteProtype) GetString() string {
	return this.s
}
