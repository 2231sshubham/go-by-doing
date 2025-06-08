package main

import (
	"fmt"
	"maps"
)

func main() {
	mp := make(map[string]string)
	mp["k1"] = "v1"
	mp["k2"] = "v2"
	fmt.Println(mp)

	delete(mp, "k1")
	_, presence := mp["k2"]
	fmt.Println("prsc", presence)

	mp1 := map[string]int{"k1": 1, "k2": 2}
	mp2 := map[string]int{"k1": 1, "k2": 2}
	if maps.Equal(mp1, mp2) {
		fmt.Println("maps equal")
	}
	for key, value := range mp1 {
		fmt.Println(key, value)
	}
}
