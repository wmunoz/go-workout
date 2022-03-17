package main

import (
	"fmt"
)

func main() {
	s1 := []int{1,3,5}
	s2 := []int{2,4,6}
	r := combine(s1, s2)
	fmt.Println(r)
}

func combine(s1 []int, s2 []int) []int {
	var s []int
	for i, _ := range s1 {
		s = append(s, s1[i])
		s = append(s, s2[i])
	}
	return s
}