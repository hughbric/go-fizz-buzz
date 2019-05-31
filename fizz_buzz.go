package main

import "fmt"

func FizzBuzz(number int) string {
	result := ""
	if number%3 == 0 {
		result += "fizz"
	} else if number%5 == 0 {
		result += "buzz"
	}
	return result
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
