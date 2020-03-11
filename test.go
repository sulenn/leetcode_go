package main

import "fmt"

type Employee struct {
	Name string
	Age int
}
func main() {
	m := map[string]*Employee{}
	workers := []Employee{
{
	"Michcle",
	30,
},
{
	"Nick",
	35,
},
}
	for _, worker := range workers {
		m[worker.Name] = &worker
}
for k, v := range m {
	fmt.Println(k)
	fmt.Println(v)
}
}