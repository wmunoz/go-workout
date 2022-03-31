package main

import (
	"fmt"
)

func main() {
	n1 := 1
	n2 := 1

	for i := 3; i <= 100; i++ {
		sum := n1 + n2
		fmt.Printf("Fibonacci number %d is %d\n", i, sum)
		n1 = n2
		n2 = sum
	}
}