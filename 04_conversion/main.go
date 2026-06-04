package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our PIzza Store")
	fmt.Println("Please the pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input);

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println("There is error ", err)
	}else{
		fmt.Printf("The datatype of numRating, %T", numRating)
	}

}
