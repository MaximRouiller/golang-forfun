package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("Testing an add function ", add(2, 4))
}
