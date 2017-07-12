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

func deferFun() {
	defer fmt.Println("Welcome to the playground!")
	defer fmt.Println("Testing an add function ", add(2, 4))
	a, b := swap("world", "hello")
	defer fmt.Println("Apparently, we love tuples? ", a, b)

	defer fmt.Println(split(17))
	defer fmt.Println(math.Max(4, 1))

	defer fmt.Println(runtime.GOOS, runtime.GOARCH)
}
