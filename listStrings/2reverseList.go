package main

import (
	"fmt"
)

func main() {
	l := []int{10,3,5,1}
	reversed := reverseList(l)
	fmt.Println(reversed)
}

func reverseList(l []int) []int {
	var tmp int
	lenght := len(l)
	for i:= 0; i < lenght/2; i++ {
		tmp = l[i]
		l[i] = l[lenght - 1 - i]
		l[lenght - 1 - i] = tmp
	}
	return l	
}