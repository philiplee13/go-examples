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

	// anon func
	go func(msg string) {
		fmt.Println(msg)
	}("going in anon func")

	// need to wait https://stackoverflow.com/questions/42217995/why-doesnt-this-go-code-print-anything-with-a-goroutine
	// or use a waitgroup
	time.Sleep(time.Second * 2)
	fmt.Println("done")

	// channels
	fmt.Println("-------------channels-----------------")
	channel := make(chan string)
	sendmessage(channel, "hello")
	sendmessage(channel, "hello2") // by default channels are unbuffered
	// this means that as messages come in -> it's like a stack
	// newest messages are on top
	// we need to pass in a size if we want this to be unbuffered
	// so something like make(chan string, 2)
	fmt.Println("sleeping for 2 seconds")
	time.Sleep(time.Second * 2)
	getmessage(channel)
	getmessage(channel)

	fmt.Println("-------------channel sync-----------------")
	boolchannel := make(chan bool)
	go worker(boolchannel)
	<-boolchannel // this will block (prevent the program from finishing) until the worker is done

	fmt.Println("------select------")
	selectexample()
}

// loop that takes in a from param and prints it with the index
func loop(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// channels
func sendmessage(channel chan<- string, message string) { // arrow is on right side bc we're sending
	// this is called channel directions
	fmt.Println("sending message", message)
	go func() { channel <- message }()
}

func getmessage(channel <-chan string) { // notice the difference in the arrows
	// arrow is on left side bc we're receiving
	msg := <-channel
	fmt.Println("got from channel:", msg)
}

// channel sync
func worker(channel chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	channel <- true
}

// select

func selectexample() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ { // this loop needs to go twice because we're waiting for
		// 2 channels
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
