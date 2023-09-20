package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var contador int64
	goRoutines := 1000
	var wg sync.WaitGroup

	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}

// para rodar, use go run -race main.go
