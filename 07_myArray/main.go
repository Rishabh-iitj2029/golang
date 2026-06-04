package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array in Golang")

	var fruitList [4]string;
	
	fruitList[0] = "banana"
	fruitList[1] = "apple"
	fruitList[3] = "orange"

	fmt.Println("The fruitList is ", fruitList)
	fmt.Println("The fruitList is ", len(fruitList))

}
