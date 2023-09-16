package main

import (
	"fmt"
	"time"
)

func main() {
	year := 1993
	currentYear := time.Now().Year()

	for {
		if year > currentYear {
			break
		}
		fmt.Println(year)
		year++
	}
}