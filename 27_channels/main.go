package main

import (
	"fmt"
	"sync"
)

var mych = make(chan int,1)

func main() {
	fmt.Println("welcome to channels")

	wg := &sync.WaitGroup{}

	// 	fmt.Println(<-ch)
	// 	ch <- 45;

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mych)
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		
		mych<-42
		mych<-34
		close(mych)
		wg.Done()
	}(mych, wg)
	wg.Wait()
}
