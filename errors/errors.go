package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(f(2))
	fmt.Println(f(-2))
}

func f(arg int) (int, error) {
	if arg < 0 {
		return 0, errors.New("Can't be used with negative numbers")
	}
	return arg + 0, nil
}
