package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main() {
	var name string
	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	name = strings.TrimRight(input, "\n")
	fmt.Println("You entered:", name)
	if (name != "Alice") && (name != "Bob") {
		name = ""
	}
	fmt.Println("Hello", name)

}