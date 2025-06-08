package main

//type List[T any] struct {
//	head, tail *element[T]
//}
//type element[T any] struct {
//	next *element[T]
//	val  T
//}
//
//func (l *List[T]) push(v T) {
//	if l.head == nil {
//		l.head = &element[T]{val: v}
//		l.tail = l.head
//	} else {
//		l.tail.next = &element[T]{val: v}
//		l.tail = l.tail.next
//	}
//}
//
//func PrintAll[T any](items []T) {
//	for _, item := range items {
//		fmt.Println(item)
//	}
//}
//func PrintAllElements[T any](ll List[T]) {
//	head := ll.head
//	for e := head; e != nil; e = e.next {
//		fmt.Println(e.val)
//	}
//}
//func main() {
//	//v := []int{1, 2, 3, 4, 5}
//	//PrintAll(v)
//
//	var ele = []int{1, 2, 3, 4}
//	var ll List[int]
//	for _, item := range ele {
//		ll.push(item)
//		//fmt.Println(item)
//	}
//	PrintAllElements(ll)
//}
