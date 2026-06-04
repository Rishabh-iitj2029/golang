package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the slices")

	var fruitList = []string{"apple", "lichi", "mango"}

	fmt.Printf("The datatype of fruitList => %T\n", fruitList)

	fruitList = append(fruitList, "papaya","watermlon")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highSocre := make([]int,4)

	highSocre[0] = 784
	highSocre[1] = 234
	highSocre[2] = 724
	highSocre[3] = 584
	// highSocre[4] = 798

	highSocre = append(highSocre, 785, 465, 654)

	fmt.Println(highSocre)

	sort.Ints(highSocre)
	fmt.Println(highSocre)

	fmt.Println("\n-------------------------------------------------------------")
	// how to remove value from the based on the index

	courses := []string{"reaactjs","expressjs","fastAPI", "python"}
	fmt.Println(courses)

	index := 2
	courses = append(courses[:index],courses[index+1:]...)

	fmt.Println(courses)
}
