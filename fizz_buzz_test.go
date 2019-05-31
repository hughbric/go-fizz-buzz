package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	fizz := FizzBuzz(3)
	if fizz != "fizz" {
		t.Errorf("Number did not fizz at %d", 3)
	}

	buzz := FizzBuzz(5)
	if buzz != "buzz" {
		t.Errorf("Number did not buzz at %d", 5)
	}

	fizzbuzz := FizzBuzz(15)
	if fizzbuzz != "fizzbuzz" {
		t.Errorf("Number did not fizzbuzz at %d", 15)
	}

	number := FizzBuzz(1)
	if number != "1" {
		t.Errorf("Number did return %d", 1)
	}

	zero := FizzBuzz(4)
	if zero != "4" {
		t.Errorf("Number did return %d", 4)
	}
}
