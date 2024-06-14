package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var result int = add(1, 2)
	fmt.Printf("result is %v\n", result)
}

// Add Method
// Example method
func add(x int, y int) int {
	var answer int = x + y
	return answer
}

// Loop Method
// Prints each index up until the limit defined
func loop(limit int) {
}
