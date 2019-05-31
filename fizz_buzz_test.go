package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	number := 3
	fizz := FizzBuzz(number)
	if fizz != `fizz` {
		t.Errorf("Number did not fizz at %d", number)
	}
}
