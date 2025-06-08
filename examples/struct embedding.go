package main

import "fmt"

type base struct {
	number int
}

func (b base) describe() string {
	return fmt.Sprintf("%v", b.number)
}

type container struct {
	base
	str string
}

func main() {
	c := container{
		base: base{
			number: 10,
		},
		str: "weird-string",
	}
	fmt.Printf("co={num = %v, str = %v}\n", c.number, c.str)
	fmt.Println("number:", c.base.number)
	fmt.Println(c.describe())
	//Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was embedded from base directly on co.

	type describer interface {
		describe() string
	}

	var d describer = c
	fmt.Println(d.describe())
	//Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a container now implements the describer interface because it embeds base
}
