package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of go lang")

	presentTime := time.Now()
	fmt.Println("The present time is => ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2006,time.June,25,23, 23,23,0,time.UTC)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
