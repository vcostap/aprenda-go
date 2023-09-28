package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go inputValues(c)

	for i := 1; i <= 100; i++ {
		fmt.Println(<-c)
	}
}

func inputValues(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}
