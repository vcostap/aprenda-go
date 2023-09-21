package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Sistema operacional:", runtime.GOOS)
	fmt.Println("Arquitetura:", runtime.GOARCH)
}
