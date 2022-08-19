package main

import (
	"fmt"
	"pointers/pointers"
)

func main() {
	fmt.Println("Pointers")

	var number int8 = 10

	pointers.CreatePointer(number)
}
