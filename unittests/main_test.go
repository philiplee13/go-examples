package unittests

import (
	"testing"
)

func TestAddMethodShouldPass(t *testing.T) {
	result := 5
	value := Add(2, 3) // because this is in the same directory, this is okay
	if result != value {
		t.Fatalf("Values did not match up, wanted %v but got %v instead", result, value)
	}
}

func TestAddMethodShouldFail(t *testing.T) {
	result := 3
	value := Add(3, 3)
	if result != value {
		t.Fatalf("Values did not match, wanted %v, but got %v instead", result, value)
	}
}
