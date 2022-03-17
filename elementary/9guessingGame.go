package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	randomNumber  := r.Intn(20)

	for {
		fmt.Println("Guess my number:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		userNumber, _ := strconv.Atoi(input)

		if userNumber > randomNumber {
			fmt.Println("Too high")
		} else if userNumber < randomNumber {
			fmt.Println("Too low")
		} else {
			fmt.Println("You won!")
			break
		}
	}
}
