package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 3; i++ {
		fmt.Println(name, "Running ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// adding go ahead of the method call makes each method call run concurrently
	// var wg sync.WaitGroup
	// wg.Add(2) // number of goroutines
	// go task("Task 1", &wg)
	// go task("task 2", &wg)
	// wg.Wait()
	// time.Sleep(time.Second * 2) // time.Sleep like this to wait for all goroutines to finish is finicky

	// doing things this way - means we can only send one value at a time
	// numberChannel := make(chan int)
	// go sendData(numberChannel, 20)
	// go getData(numberChannel)
	// go sendData(numberChannel, 30)
	// go getData(numberChannel)

	bufferedChannel := make(chan int, 5)
	go sendData(bufferedChannel, 1)
	go sendData(bufferedChannel, 2)
	go sendData(bufferedChannel, 3)
	go getData(bufferedChannel)
	go getData(bufferedChannel)
	go getData(bufferedChannel)

	time.Sleep(time.Second * 5)
	fmt.Println("All goroutines finished")
}

// channels and it's usage
func sendData(channel chan<- int, number int) {
	fmt.Println("Sending number", number)
	channel <- number
	fmt.Println("Finished sending", number)
}

func getData(channel <-chan int) {
	value := <-channel
	fmt.Println("Got value", value)
}
