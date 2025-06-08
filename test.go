package main

import "fmt"

func willPanic() {
	panic("boom")
}

func recoverLater() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Too late to recover here")
		}
	}()
	willPanic() // this crashes before recover can be called
}

func main() {
	recoverLater() // program crashes unless recovered in `willPanic`
}

// func inner() string {
// 	panic("inner panic")
// }

// func safeDivide() (res string) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered:", r)
// 			// res = "recovered from panic"
// 		}
// 	}()
// 	res = inner()
// 	return "returned"
// }

// func main() {
// 	fmt.Println(safeDivide()) // ✅ prints -1
// }
