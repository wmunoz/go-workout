package main

import (
	"fmt"
)

func main() {
	A := []int{1,2,3,4,5,6}
	rotateSlice(&A, 2)
	fmt.Println(A)
}

func rotateSlice(arr *[]int, n int) {
	l := len(*arr)
	for i := 1; i <= n; i++ {
		*arr = append(*arr, 0)
	}
	copy((*arr)[l:], (*arr)[:n])
	(*arr) = (*arr)[n:]
}