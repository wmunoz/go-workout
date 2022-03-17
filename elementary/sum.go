package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func computeSum(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i
	}	
		return sum
}

func computeProduct(number int) int {
	product := 1
	for i := 1; i <= number; i++ {
		product *= i
	}	
		return  product
}

func main() {
	fmt.Println("Please enter a number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")

	fmt.Println("Please enter and operation:")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimRight(operation, "\n")

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Errorf("An error ocurred: %v", err)
	}
	var result int
	switch operation {
	case "sum":
		result = computeSum(number)
	case "product":
		result = computeProduct(number)
	default:
		result = 0
	}
	fmt.Println("The result is:", result)
}

