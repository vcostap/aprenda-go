package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Iniciando o programa. Número de goroutines:", runtime.NumGoroutine())

	c := make(chan int)

	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				c <- j
			}
		}()
		fmt.Println("Número de goroutines:", runtime.NumGoroutine())
	}

	for i := 1; i <= 100; i++ {
		fmt.Println("i:", i, "chan value:", <-c)
	}
	fmt.Println("Finalizando o programa. Número de goroutines:", runtime.NumGoroutine())
}
