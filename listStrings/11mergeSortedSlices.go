package main

import (
	"fmt"
)

func main() {
 b := []int{8,20}
 a := []int{1,5,9}
 result := mergeSortedSlices(a, b)
 fmt.Println(result)

}

func mergeSortedSlices(A []int, B []int) []int {
	positionTracker := 0
	for _, a := range A {
		inserted := false
		for j := positionTracker; j < len(B); j++ {
			if a < B[j] { 
				insertIntoPosition(&B, j, a)
				positionTracker = j + 1
				inserted = true
				break
			}
		}
		if !inserted {
			B = append(B, a)
		}
	}
	return B
}

func insertIntoPosition(arr *[]int, pos int, newElement int)  {
	*arr = append(*arr, 0)
	copy((*arr)[pos+1:], (*arr)[pos:])
	(*arr)[pos] = newElement
}