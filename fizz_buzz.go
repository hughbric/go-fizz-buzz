package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(number int) string {
	result := ""
	if number%3 == 0 {
		result += "fizz"
	}
	if number%5 == 0 {
		result += "buzz"
	}
	if result == "" {
		result += strconv.Itoa(number)
	}
	return result
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
