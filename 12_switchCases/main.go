package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the switch cases in the Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(5) + 1

	fmt.Println("the dice number is => ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Printf("The dice value is %v and no more chances\n", diceNumber)

	case 2:
		fmt.Printf("The dice value is %v and no more chances\n", diceNumber)
		fallthrough
	case 3:
		fmt.Printf("The dice value is %v and no more chances\n", diceNumber)

	case 4:
		fmt.Printf("The dice value is %v and no more chances\n", diceNumber)

	case 5:
		fmt.Printf("The dice value is %v and no more chances\n", diceNumber)

	case 6:
		fmt.Printf("The dice value is %v and roll again\n", diceNumber)
	

	default:
		fmt.Println("wrong input")
	}
}
