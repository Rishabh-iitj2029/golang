package main

import (
	"fmt"
	"net/http"
	"sync"
)

var counter int

func main() {

	fmt.Println("---------------SENDING REQUEST------------------")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	urlendpoints := []string{
		"https://www.google.com/",
		"https://projects.devluplabs.tech/projects",
		"https://github.com/Rishabh-iitj2029",
		"https://leetcode.com/",
		"https://medium.com/@shikha.ritu17/rest-api-architecture-6f1c3c99f0d3",
		"https://www.youtube.com/watch?v=i5HEE5Ikv3w&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=56",
		"https://pkg.go.dev/",
	}

	wg.Add(len(urlendpoints))
	for ind, url := range urlendpoints {
		go func(wg *sync.WaitGroup, mut *sync.Mutex) {
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}

			defer wg.Done()

			if res.StatusCode == 200 {
				mut.Lock()
				counter++
				mut.Unlock()
			}
			fmt.Printf("%v index -- url %v -- status code %v\n", ind+1, url, res.StatusCode)

		}(wg,mut)
	}
	wg.Wait()

	fmt.Println("The counter value = ", counter)
}
