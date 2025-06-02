package main

import "fmt"

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
} // Function demonstrating Closure. We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt

func main() {
	f := initSeq()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	f = initSeq()
	fmt.Println(f())

	//Anonymous functions can also be recursive, but this requires explicitly declaring a variable with var to store the function before it’s defined
	var fib func(int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n) + fib(n-2)
	} // Recursive function
}
