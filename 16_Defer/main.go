package main

import "fmt"


func main()  {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Welcome to the defer")

	myDefer(10)

}

func myDefer(a int){
	for i := 0; i < a; i++ {
		defer fmt.Println("The value is ", i);
	}
}