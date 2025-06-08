package main

import "fmt"

func main() {
	var slice []string
	fmt.Println(slice, len(slice), cap(slice))

	slice = make([]string, 4, 5)
	fmt.Println(slice, len(slice), cap(slice))

	slice = []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(slice)

	n := slice[:3]
	fmt.Println(n)
	n1 := slice[1:4]
	fmt.Println(n1)
	n2 := slice[:2]
	fmt.Println(n2)

	n3 := make([]string, 6)
	copy(n3, slice) // deep copy

	n1[2] = "h"
	fmt.Println(n)
	fmt.Println(slice)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3) // it doesn't have any affect of n1[2]="h" statement as n3 was deep copied using copy()

}
