package main

import "testing"

func TestFizzbuzzReturnNumberIfNotDivisible(t *testing.T) {
	actual := Fizzbuzz(7)
	expected := "7"

	if actual != expected {
		t.Errorf("actual '%s' expected '%s'", actual, expected)
	}
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByThree(t *testing.T) {
	actual := Fizzbuzz(3)
	expected := "Fizz"

	if actual != expected {
		t.Errorf("actual '%s' expected '%s'", actual, expected)
	}
}

func TestFizzbuzzKnowsIfANumberIsNotDivisibleByThree(t *testing.T) {
	actual := Fizzbuzz(4)
	expected := "4"

	if actual != expected {
		t.Errorf("actual '%s' expected '%s'", actual, expected)
	}
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByFive(t *testing.T) {
	actual := Fizzbuzz(5)
	expected := "Buzz"

	if actual != expected {
		t.Errorf("actual '%s' expected '%s'", actual, expected)
	}
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByFifteen(t *testing.T) {
	actual := Fizzbuzz(15)
	expected := "Fizzbuzz"

	if actual != expected {
		t.Errorf("actual '%s' expected '%s'", actual, expected)
	}
}
