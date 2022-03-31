package main

import (
	"fmt"
)

func main() {
	n := 12345
	mod := 1
	A := []int{}
	for n > 0 {
		mod = n % 10
		n = n / 10
		A = prependInt(A, mod)
	}
	fmt.Println(A)
}

func prependInt(arr []int, n int) []int {
	arr = append(arr, 0)
	copy(arr[1:], arr)
	arr[0] = n
	return arr
}