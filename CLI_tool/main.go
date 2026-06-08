package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)



type responseType struct {
	Endpoint   string
	StatusCode int
}

func main() {

	for {
		var counter int64

		fmt.Println("---------------SENDING REQUEST------------------")
		wg := &sync.WaitGroup{}

		urlendpoints := []string{
			"https://www.google.com/",
			"https://projects.devluplabs.tech/projects",
			"https://github.com/Rishabh-iitj2029",
			"https://leetcode.com/",
			"https://medium.com/@shikha.ritu17/rest-api-architecture-6f1c3c99f0d3",
			"https://www.youtube.com/watch?v=i5HEE5Ikv3w&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=56",
			"https://pkg.go.dev/",
		}

		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		ch := make(chan responseType, len(urlendpoints))

		for _, url := range urlendpoints {
			wg.Add(1)
			go func(target string) {
				defer wg.Done()
				res, err := client.Get(target)
				if err != nil {
					fmt.Println(err)
				}

				if res.StatusCode == 200 {
					atomic.AddInt64(&counter, 1)
				}

				ch <- responseType{target, res.StatusCode}
				res.Body.Close()

			}(url)
		}

		results := make([]responseType,0,len(urlendpoints))

		go func() {
			wg.Wait()
			close(ch)
		}()

		for url := range ch {
			results = append(results, url)
		}

		for _, val := range results {
			fmt.Printf("%v\n", val)
		}
		fmt.Println("The counter value = ", counter)
		fmt.Printf("\n")
		time.Sleep(2 * time.Second)
	}
}
