package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in go lang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// for d:= 0; d < len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	for _, value := range days {
		fmt.Printf("The index is and value is %v\n", value)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		if rougueValue == 4 {
			goto hi
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

hi:
	fmt.Println("It's hi from rishabh")
}
