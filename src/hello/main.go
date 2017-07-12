package main

import (
	"fmt"
)

type GeoPoint struct {
	X float32
	Y float32
}

func main() {
	fmt.Println("first")
	deferFun()
	playingWithStruct()
	fmt.Println("last")
}
