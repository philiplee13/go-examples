package main

import "fmt"

func main() {
	genericmethod()
	fmt.Println(singlereturntype())
	fmt.Println(multiplereturntypes(false))
	fmt.Println(multiplereturntypes(true))
	varadicfunction(1, 2)
	varadicfunction(1, 2, 3)

	numanonfunc := anonfunc()
	fmt.Println(numanonfunc())
	fmt.Println(numanonfunc()) // the reason why this increments is bc
	// we assigned the actual function to the var
	// and we're continously calling it
	// so for each call -> the newly incremented number is getting returned
}

func genericmethod() {
	fmt.Println("generic method no return type")
}

func singlereturntype() string {
	return "single return type"
}

func multiplereturntypes(triggererror bool) (string, error) {
	if triggererror {
		return "", fmt.Errorf("error occurred in multiple return types")
	}
	return "multiple return type no errors", nil
}

func varadicfunction(numbers ...int) int {
	total := 0
	for idx, num := range numbers {
		fmt.Println("index:", idx, "value:", num)
		total += num
	}
	return total
}

func anonfunc() func() int {
	fmt.Println("start of anon func")
	number := 0
	fmt.Println("number at start is", number)
	return func() int {
		number++
		fmt.Println("number inside of anon func is ", number)
		return number
	}
}
