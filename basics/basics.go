package main

import (
	"fmt"

	"main/numbers"
)

// Strucs
type Human struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello World")
	loop(5)
	arr := []string{"hello", "my", "name", "is", "dog"}
	rangeLoop(arr)

	hashmap := map[string]string{"keyOne": "valueOne", "keyTwo": "valueTwo"}
	iterateThroughMap(hashmap)

	jane := Human{name: "Jane", age: 10}
	printDetails(jane)
	fmt.Println("calling the add method from numbers", numbers.Add(2, 3))

	fmt.Println("calling subtract method from numbers", numbers.Subtract(3, 1))
}

// Loop Method
// Prints each index up until the limit defined
func loop(limit int) {
	// loop using for x ... syntax
	for i := 0; i < limit; i++ {
		fmt.Printf("i is %v\n", i)
	}
}

// Range Loop Method
// Prints the index and element of each element in the array
func rangeLoop(arr []string) {
	for i := range arr {
		fmt.Printf("index is %v\n element is %v\n", i, arr[i])
	}
}

// Iterates through a hashmap
// Prints out the key and value for each element in map
func iterateThroughMap(hashmap map[string]string) {
	for k, v := range hashmap {
		fmt.Printf("key is %v, and element is %v\n", k, v)
	}
}

func printDetails(human Human) {
	fmt.Printf("name is %v\n", human.name)
	fmt.Printf("age is %v\n", human.age)
}
