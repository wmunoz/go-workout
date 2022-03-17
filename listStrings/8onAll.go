package main

import (
	"fmt"
)

func main() {
	numbers := []int{}
	for i := 1; i <= 20; i++ {
		numbers = append(numbers, i)
	} 
	r := onAll(numbers)
	fmt.Println(r)
}

func onAll(s []int) []int {
	result := []int{}
	for _, n := range s {
		result = append(result, n*n)
	}
	return result
}