package main

import "fmt"

func main() {
	var a [4]int
	a[1] = 12
	fmt.Println(a)

	a = [4]int{1, 2, 3, 4}
	fmt.Println(a)

	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	fmt.Println("Length:", len(b))

	b = [...]int{1, 3: 3, 4}
	fmt.Println(b)

	var arr = [2][3]string{
		{"asb", "Sa", "fasf"}, {"Faf", "faf", "fae"},
	}
	fmt.Println(arr)
}
