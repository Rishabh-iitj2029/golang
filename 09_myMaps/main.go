package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in goLang")

	languages := make(map[string]string)

	languages["Js"] = "Javascript"
	languages["py"] = "Python"
	languages["cpp"] = "C++"

	fmt.Println("List of all languages", languages)

	fmt.Println("Js is shortform for ", languages["Js"])

	delete(languages,"cpp")
	fmt.Println("List of all languages", languages)

	// loops in maps

	for key, value := range languages{
		fmt.Printf("The shortform of %v : %v\n",key, value )
	}

	for _, value := range languages{
		fmt.Printf("The shortform of  : %v\n", value )
	}
}