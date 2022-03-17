package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	year := t.Year()
	leapYear := true

	for y := year; y <= year + 20; y++ {
		if y % 4 != 0 {
			leapYear = false
		} else if y % 100 != 0 {
			leapYear = true
		} else if y % 400 != 0 {
			leapYear = false
		} else {
			leapYear = true
		}
		if leapYear == true {
			fmt.Println(y)
		}
	}
}