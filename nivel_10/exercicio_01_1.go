package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//go routine 2 - função anônima
	go func() {
		c <- 42
	}()

	//go routine 1 - main
	fmt.Println(<-c)
}
