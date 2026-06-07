package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urlendpoint = []string{
	"https://www.google.com/",
	"https://www.github.com/",
	"https://www.youtube.com/",
	"https://go.dev/",
	"https://web.whatsapp.com/",
}

var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// go greeter("hello")
	// greeter("world")

	for _, value := range urlendpoint {

		go getStatusCode(value)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string){
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Printf("%v = %v\n", i,s);
// 	}
// }

func getStatusCode(url string) {

	result, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		mut.Lock()
		signals = append(signals, url)
		mut.Unlock()
		fmt.Printf("The Status Code for %v is %v\n", url, result.StatusCode)
		defer wg.Done()
	}
}
