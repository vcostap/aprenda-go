package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Goroutines antes:", runtime.NumGoroutine())
	c := gen()
	fmt.Println("Goroutines depois:", runtime.NumGoroutine())
	receive(c)
	fmt.Println("Goroutines no fim:", runtime.NumGoroutine())
	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Println("Goroutines durante:", runtime.NumGoroutine())
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
