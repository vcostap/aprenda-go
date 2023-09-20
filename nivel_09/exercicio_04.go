package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	contador := 0
	goRoutines := 1000
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}

// para rodar, use go run -race main.go
