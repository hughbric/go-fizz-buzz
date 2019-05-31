package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(number int) string {
	result := ""
	if number%3 == 0 && number%5 == 0 {
		result += "fizzbuzz"
	} else if number%3 == 0 {
		result += "fizz"
	} else if number%5 == 0 {
		result += "buzz"
	} else {
		result += strconv.Itoa(number)
	}
	return result
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
