package main

import "fmt"

func main() {
	fmt.Println("Welcome to the functions in Go")
	greeting()

	result := adder(7, 5)
	fmt.Println(result)

	total_sum, _ := proAdder(2,3,4,5,7)
	fmt.Println(total_sum)
}

func greeting()  {
	fmt.Println("Namaste, kaise ho")
}

func adder(a int,b int)int {
	return (a+b);
}

func proAdder(values ...int)(int,string) {
	total := 0;

	for _, val := range values{
		total += val;
	}
	return  total, "This is result";
}
