package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of the delivered pizza ")

	// comma ok || err ok
	input, _:= reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	// fmt.Println("The error found is , ", err)
	fmt.Printf("The type of input is %T\n", input);

}
