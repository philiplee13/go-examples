package helpers

import (
	"testing"
)

func TestSubtractMethodShouldPass(t *testing.T) {
	result := 0
	value := Subtract(2, 2)
	if result != value {
		t.Fatalf("Values did not match up, wanted %v, but got %v instead", result, value)
	}
	t.Logf("Test passed")
}

func TestSubtractMethodShouldFail(t *testing.T) {
	result := 0
	value := Subtract(2, 1)
	if result != value {
		t.Fatalf("Values did not match up, wanted %v, but got %v instead", result, value)
	}
	t.Logf("Test passed")
}
