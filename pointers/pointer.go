package pointers

import "fmt"

func CreatePointer(n1 int8) (int8, *int8) {
	var number int8
	var pointer *int8 // Declaração de um ponteiro

	number = n1
	pointer = &number // Referenciar

	fmt.Println(number)
	fmt.Println(pointer)
	fmt.Println(*pointer)

	number++
	*pointer++

	fmt.Println(number)
	fmt.Println(pointer)  // Devolve a posição de memória
	fmt.Println(*pointer) // Desreferenciar

	return number, pointer
}
