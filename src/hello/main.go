package main

import (
	"fmt"
	"math"
	"runtime"
)

// because I need a function for that.
func add(x, y int) int {
	return x + y
}

// Hello tuples? Nice built-in feature.
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("Testing an add function ", add(2, 4))
	a, b := swap("world", "hello")
	fmt.Println("Apparently, we love tuples? ", a, b)

	fmt.Println(split(17))
	fmt.Println(math.Max(4, 1))

	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
