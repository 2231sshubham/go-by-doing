package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}
func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

//Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.

func main() {
	r := rect{width: 10, height: 5}
	n := &r
	//Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", n.perimeter())
	//Here we call the 2 methods defined for our struct.
}
