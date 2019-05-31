package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	fizz := FizzBuzz(3)
	if fizz != `fizz` {
		t.Errorf("Number did not fizz at %d", 3)
	}

	buzz := FizzBuzz(5)
	if buzz != `buzz` {
		t.Errorf("Number did not buzz at %d", 5)
	}
}
