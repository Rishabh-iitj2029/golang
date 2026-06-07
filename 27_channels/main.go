package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to channels")

	// 	fmt.Println(<-ch)
	// 	ch <- 45;
	var mych1 = make(chan int, 1)
	mychan2 := make(chan int)

	go func() {
		mych1 <- 42
		mychan2 <- 56
	}()

	res1 := <-mych1
	res2 := <-mychan2

	fmt.Printf("%v and %v\n", res1, res2)

}
