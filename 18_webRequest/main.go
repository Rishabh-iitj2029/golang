package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://api.github.com/users/Rishabh-iitj2029"

func main() {
	fmt.Println("Making a web request")
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("The respnse type is %T\n", res)

	defer res.Body.Close() //remember

	dataByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The datatype is %T\n", dataByte)
	var contentJson map[string]any

	err1 := json.Unmarshal(dataByte, &contentJson)

	if(err1 != nil){
		panic(err1)
	}
	fmt.Println("the json data is =>\n", contentJson["bio"])

	content := string(dataByte)

	fmt.Printf("The value is %v\n", content)
}
