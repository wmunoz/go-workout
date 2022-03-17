package main

import (
	"fmt"
)

func main() {
	myString := "ab"
	if isPalindrome(myString) {
		fmt.Printf("%s is palindrome!", myString)
	} else {
		fmt.Println("Is NOT")
	}
}

func isPalindrome(w string) bool {
	for i := 0; i < len(w)/2; i++ {
		if w[i] != w[len(w) - 1 - i] {
			return false
		}
	}
	return true
}