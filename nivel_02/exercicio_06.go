package main

import "fmt"

const (
	years = 2023 + iota*4
	years4
	years8
	years12
)

func main() {
	fmt.Println(years, years4, years8, years12)
}