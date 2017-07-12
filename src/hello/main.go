package main

import (
	"fmt"
)

func ranges() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	for i := range pow {
		fmt.Printf("%d\n", i)
	}
}
func main() {
	fmt.Println("first")
	//deferFun()
	//playingWithStruct()
	ranges()
	fmt.Println("last")
}
