package main

import "fmt"

func FizzBuzz(number int) string {
	result := ""
	if number%3 == 0 {
		result += "fizz"
	}
	return result
}

func main() {
	a := FizzBuzz(3)
	fmt.Println(a)
}
