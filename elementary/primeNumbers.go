package main

import (
	"fmt"
	"math"
)

func main() {
	primes := []int{2}
	for i := 3; i <= math.MaxInt; i ++ {
		isPrime := true
		for _, j := range primes {
			if i % j == 0 {
				fmt.Println("Not prime", i)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Println("Prime", i)
			primes = append(primes, i)
		}
	} 
}

