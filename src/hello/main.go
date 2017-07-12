package main

import (
	"fmt"
)

/// test
/* testing some more*/
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("Testing an add function ", add(2, 4))
}
