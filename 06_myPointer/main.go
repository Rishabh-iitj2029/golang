package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of pointer")

	// var ptr *int;
	// fmt.Println(ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of the pointer is => ", ptr)
	fmt.Println("Value stored in the pointer =>", *ptr)

	*ptr = *ptr * 2;

	fmt.Println("The new value is => ", myNumber)
}
