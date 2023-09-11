package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("printei primeiro")
	fmt.Println("printei segundo")
}