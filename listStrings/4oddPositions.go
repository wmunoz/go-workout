package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"a", "b", "c", "d"}
	result := findOddPositionedElement(mySlice)
	fmt.Println(result)
}

func findOddPositionedElement(l []string) []string {
	result := []string{}
	for index, value := range(l) {
		if index % 2 != 0 {
			result = append(result, value)
		}
	}
	return result
}