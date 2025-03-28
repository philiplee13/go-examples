package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("sequential run")
	loop("sequential")

	fmt.Println("starting goroutines")
	go loop("routine1")
	go loop("routine2")

	// need to wait https://stackoverflow.com/questions/42217995/why-doesnt-this-go-code-print-anything-with-a-goroutine
	time.Sleep(time.Second * 2)
	fmt.Println("done")
}

func loop(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
