package main

import "fmt"

func main() {
	var prototype Prototype = new(ConcreteProtype)
	prototype.Construct("Clone")
	var clone Prototype = prototype.Clone()
	fmt.Println(prototype.GetString())
	fmt.Println(clone.GetString())
}
