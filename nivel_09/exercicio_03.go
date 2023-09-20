package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	contador := 0
	goRoutines := 1000

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(contador)
}

// para rodar, use go run -race main.go
