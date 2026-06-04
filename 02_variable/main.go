package main

import "fmt"

var jwtToken int = 6514654981
const RefreshToken string = "awbwerywrsr7426947greg" 

func main() {
	var username string = "Rishabh"
	fmt.Println(username)
	fmt.Printf("Variable type is : %T\n", username)

	var isVerified bool = false
	fmt.Println(isVerified)
	fmt.Printf("Variable type is : %T\n", isVerified)

	var smallValue int = 256
	fmt.Println("The value is ",smallValue)
	fmt.Printf("Variable type is : %T\n", smallValue)

	var samllFloat float64 = 255.457878142472147
	fmt.Println(samllFloat)
	fmt.Printf("Variable type is : %T\n", samllFloat)

	var another string;
	fmt.Println(another)
	fmt.Printf("The type of variable is %T\n", another)

	// Implicit type of declaring the var

	var website = "hello it's my portfolio"
	fmt.Println(website)

	// no var style

	numberOfUser := 4578
	fmt.Println(numberOfUser)
	fmt.Println(jwtToken)


	fmt.Println(RefreshToken)
	fmt.Printf("The type of variable is %T\n", RefreshToken)
}
