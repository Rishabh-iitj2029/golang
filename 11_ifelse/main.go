package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the IfElse in golang")

	logInCount := 23

	var result string
	if logInCount < 10 {
		result = "Regular User"
	} else if logInCount > 10 && logInCount <= 23 {
		result = "something else"
	} else {
		result = "something went wrong"
	}

	fmt.Println(result)

	var Age int;
	fmt.Println("Enter your age ")
	fmt.Scanf("%d", &Age)
	if(Age>=18){
		fmt.Println("You are adult")
	}else{
		fmt.Println("You are a teenager")
	}


	if num := 5; num < 10{
		fmt.Println("Authenciated")
	}else{
		fmt.Println("Unauthorized access")
	}

	
}
