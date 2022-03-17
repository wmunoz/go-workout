package main

import (
	"fmt"
)

func main() {
	s1 := []int{1,2}
	s2 := []int{3,4}

	concatenate(&s1, s2)
	fmt.Println(s1)
}

func concatenate(s1 *[]int, s2[]int) {
	*s1 = append(*s1, s2...)
}