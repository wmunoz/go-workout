package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("What multiplication table do you want to print out? (numbers up to 12)")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")
	number, _ := strconv.Atoi(input)
	for i := 1; i <= 12; i++ {
		product := number * i
		fmt.Printf("%d x %d = %d\n", number, i, product)
	}
}