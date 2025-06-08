package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name, age: 42}
	return &p
}

// Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it
func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 42})
	fmt.Println(person{name: "Bob"})
	fmt.Println(&person{name: "Ann", age: 35}) //An & prefix yields a pointer to the struct
	fmt.Println(newPerson("John"))             //It’s idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Dennis", age: 50}
	n := &s
	fmt.Println(n.age) // Go auto-deference the pointer variable

	dog := struct {
		name   string
		isGood bool
	}{
		"Freddy",
		true,
	}
	fmt.Println(dog) //If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests
}
