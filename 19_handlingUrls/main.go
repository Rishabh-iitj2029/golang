package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.google.com/search?q=converting+bytes+to+json+in+go&ie=UTF-8&course=golangInEnglish&paymentId=74582ievsaqaf"

func main() {
	fmt.Println("Handling urls in Go")

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is => %T\n", qparams)
	fmt.Println(qparams)
	fmt.Println(qparams["paymentId"])

	for _, value := range qparams {
		fmt.Printf("The query is %v\n", value)
	}

	partOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "www.google.com",
		Path:    "/search",
		RawPath: "user=Rishabh",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println(anotherUrl)
}
