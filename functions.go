package main

import "fmt"

// Functions are central in Go

func sum(a, b int) int {
	return a + b
}

func val() (int, int) {
	return 4, 5
} // multple return value function

// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
func variadicSum(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
} // variadic function

func main() {
	fmt.Println(sum(1, 2))

	a, b := val()
	fmt.Println(a, b)

	_, c := val()
	fmt.Println(c)

	slice := []int{1, 2, 3}
	fmt.Println(variadicSum(slice...))
	fmt.Println(variadicSum(1, 4, 5))
}
