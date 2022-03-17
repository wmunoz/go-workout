package main

import (
	"fmt"
)

func main() {
	l := []int{10,3,5,7,12,8}
	max := largestElement(l)
	fmt.Println(max)
}

func largestElement(l []int) int {
	max := l[0]
	for i:=1; i < len(l); i++ {
		if l[i] > max {
			max = l[i]
		}
	}
	return max
}