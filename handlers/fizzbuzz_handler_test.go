package handlers

import (
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	begin, _ := FizzBuzz(3, 5, 10, "fizz", "buzz")
	end := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz"}
	if strings.Join(begin, ",") != strings.Join(end, ",") {
		t.Errorf("FizzBuzz \n got: \n%v \n want: \n%v", begin, end)
	}
}

func TestFizzBuzzWithNegativeLimit(t *testing.T) {
	_, err := FizzBuzz(3, 5, -10, "fizz", "buzz")
	expected := "limit parameter must be positive"
	if expected != err.Error() {
		t.Errorf("Error actual = %v, and Expected = %v.", err.Error(), expected)
	}
}

func TestFizzBuzzWithZeroLimit(t *testing.T) {
	_, err := FizzBuzz(3, 5, 0, "fizz", "buzz")
	expected := "limit parameter must be positive"
	if expected != err.Error() {
		t.Errorf("Error actual = %v, and Expected = %v.", err.Error(), expected)
	}
}

func TestFizzBuzzWithZeroMultiple(t *testing.T) {
	_, err := FizzBuzz(0, 0, 100, "fizz", "buzz")
	expected := "multiples must be positive"
	if expected != err.Error() {
		t.Errorf("Error actual = %v, and Expected = %v.", err.Error(), expected)
	}
}
