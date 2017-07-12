package main

import (
	"fmt"
)

// because I need a function for that.
func add(x, y int) int {
	return x + y
}

// Hello tuples? Nice built-in feature.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("Testing an add function ", add(2, 4))
	a, b := swap("world", "hello")
	fmt.Println("Apparently, we love tuples? ", a, b)
}
