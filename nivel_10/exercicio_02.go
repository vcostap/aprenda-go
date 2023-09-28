package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go sender(cs, 42)

	receiver(cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func sender(s chan<- int, num int) {
	s <- num
}

func receiver(r <-chan int) {
	fmt.Println(<-r)
}
