package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rec struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rec) area() float64 {
	return r.width * r.height
}
func (r rec) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle:", c)
	} //Sometimes it’s useful to know the runtime type of an interface value. One option is using a type assertion as shown here
}
func main() {
	r := rec{10, 10}
	c := circle{10}
	measure(r)
	measure(c) //The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
}
