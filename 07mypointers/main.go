package main

import "fmt"

func main() {
	//pointers are used to store exact reference (direct reference to the memory address) or value (original not copy)
	// * is the symbol that this data type is of pointer
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)
	//* creates a pointer but & creates a pointer and referencing some memory
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("valur of ptr: ", ptr)
	fmt.Println("valur of ptr: ", *ptr)
	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)

}
