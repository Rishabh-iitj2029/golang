package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the file handling in go lang")
	content := "This is the sensitive data which i need to store"

	file, err := os.Create("./sensitivefile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println(length)
	defer file.Close()

	readFile("./sensitivefile.txt")

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("The data in the file is => \n", dataByte)
	fmt.Println("The data in the file is => \n", string(dataByte))

}

func checkNilErr(err error){
	if(err != nil){
		panic(err);
	}
}
