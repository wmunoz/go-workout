package main

import (
	"fmt"
)

func main() {
	element := "x"
	mySlice := []string{"b", "c", "d", "a"}
	if findElement(element, mySlice) {
		fmt.Printf("%s was found in the slice!\n", element)
	} else {
		fmt.Printf("%s was NOT found in the slice\n", element)
	}
}

func findElement(e string, l []string) bool {
	for _, value := range(l) {
		if e == value {
			return true
		}
	}
	return false
}