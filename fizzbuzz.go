package main

import "strconv"

func Fizzbuzz(number int) string {
	if IsDivisibleByFifteen(number) {
		return "Fizzbuzz"
	} else if IsDivisibleByFive(number) {
			return "Buzz"
	} else if IsDivisibleByThree(number) {
		return "Fizz"
	} else {
		return strconv.Itoa(number)
	}
}

func IsDivisibleByThree(number int) bool {
	return number % 3 == 0
}

func IsDivisibleByFive(number int) bool {
	return number % 5 == 0
}

func IsDivisibleByFifteen(number int) bool {
	return number % 15 == 0
}
