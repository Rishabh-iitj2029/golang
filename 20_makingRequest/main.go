package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the different web request section ")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormData()
}

func PerformGetRequest() {
	const myURL = "http://localhost:8000/get"

	res, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	dataByte, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(dataByte))

	defer res.Body.Close()

	var responseString strings.Builder

	byteCount, _ := responseString.Write(dataByte)
	fmt.Println("The content is => ", responseString.String())
	fmt.Println(byteCount)

}

func PerformPostJsonRequest() {
	const url = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"cousrsename":"Wemcome to the neural risht",
		"isVerified":"false"
	}
	`)

	res, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	dataByte, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("The content is \n", string(dataByte))

}

func PerformPostFormData() {
	myUrl := "http://localhost:8000/postform"

	//form payload

	data := url.Values{}
	data.Add("firstname", "Rishabh")
	data.Add("lastname", "Raj")
	data.Add("email", "rishabh@go.dev")

	res, err := http.PostForm(myUrl, data)

	if(err != nil){
		panic(err)
	}
	defer res.Body.Close()
	content,_ := io.ReadAll(res.Body)
	fmt.Println("The formdata response => ", string(content)) 
}
