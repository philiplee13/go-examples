package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// fmt.Println(add(2, 5))
	// loopStuff(5)
	// createDataStrucs()
	// successfulValue, valueResult := testErrHandling(5, 10)
	// failedValue, failedResult := testErrHandling(102, 105)

	// fmt.Printf("Value of successful method was %v, %v\n", successfulValue, valueResult)
	// fmt.Printf("Value of failed method was %v, %v\n", failedValue, failedResult)

	// if failedResult == false || valueResult == false {
	// 	fmt.Printf("Error happened during one of the method calls\n")
	// }
	// structMethods()
	sampleGET()
}

func add(x int, y int) int {
	var answer int = x + y
	return answer
}

func loopStuff(x int) {
	// regular for loop
	for i := 0; i < x; i++ {
		fmt.Printf("i is %v\n", i)
	}
}

func createDataStrucs() {
	// create an fixed size array
	var arr [10]int
	fmt.Println(arr)
	arr[2] = 10
	fmt.Println(arr)

	// create a dynamic sized array
	// dynamicArr := []int{10}
	var dynamicArr []int
	fmt.Println(dynamicArr)
	dynamicArr = append(dynamicArr, 100)
	fmt.Println(dynamicArr)
	dynamicArr = append(dynamicArr, 200)
	fmt.Println(dynamicArr)

	// create hash map
	testMap := make(map[string]int)
	testMap["hello"] = 1
	fmt.Println(testMap)
	testMap["bye"] = 0
	fmt.Println(testMap)
	fmt.Println(testMap["hello"])

	// loop with range over data strucs
	fmt.Println("now looping through the array with range")
	for idx, num := range arr {
		fmt.Printf("index is %v, num is %v\n", idx, num)
		fmt.Println()
	}

	fmt.Println("Now looping through the map with range")
	for k, v := range testMap {
		fmt.Printf("key is %v, value is %v\n", k, v)
		fmt.Println()
	}
}

func testErrHandling(x int, y int) (int, bool) {
	if x < 100 && y < 100 {
		return x + y, true
	}
	fmt.Println(errors.New("Error happened in testErrHandling method - either x or y was not under 100"))
	return -1, false
}

func structMethods() {
	type person struct {
		name string
		age  int
	}

	personA := person{name: "Jane Doe", age: 50}
	fmt.Println(personA)
	fmt.Printf("Name of person a is %v\n", personA.name)
	personA.name = "Bob Smith"
	fmt.Printf("Name of person a is now %v\n", personA.name)
}

func sampleGET() {
	const baseUrl string = "https://httpbin.org"
	getUrl := fmt.Sprintf("%v/get", baseUrl)
	response, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("Error happened during the get request")
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error happened when trying to read response body")
		fmt.Println(err)
	}
	fmt.Println(string(body))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
}
