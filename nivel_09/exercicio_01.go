package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go print1()
	go print2()
	wg.Wait()
}

func print1() {
	fmt.Println("Essa é a função 1")
	wg.Done()
}

func print2() {
	fmt.Println("Essa é a função 2")
	wg.Done()
}
