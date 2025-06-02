package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (l *List[T]) push(v T) {
	if l.head == nil {
		l.head = &element[T]{val: v}
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{val: v}
		l.tail = l.tail.next
	}
}
func (ll *List[T]) All() iter.Seq[T] {
	return func(express func(T) bool) {
		for e := ll.head; e != nil; e = e.next {
			if !express(e.val) {
				return
			}
		}
	}
}

// The iterator function takes another function as a parameter, called yield by convention (but the name can be arbitrary). It will call yield for every element we want to iterate over, and note yield’s return value for a potential early termination.
func genFib() iter.Seq[int] {
	return func(express func(int) bool) {
		a, b := 1, 1
		for {
			if !express(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

// Iteration doesn’t require an underlying data structure, and doesn’t even have to be finite! Here’s a function returning an iterator over Fibonacci numbers: it keeps running as long as yield keeps returning true.

func main() {
	ll := List[int]{}
	ll.push(1)
	ll.push(2)
	ll.push(3)
	for i := range ll.All() {
		fmt.Println(i)
	}
	all := slices.Collect(ll.All())
	fmt.Println("All:", all) //Packages like slices have a number of useful functions to work with iterators. For example, Collect takes any iterator and collects all its values into a slice.
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
