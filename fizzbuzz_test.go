package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestFizzbuzzReturnNumberIfNotDivisible(t *testing.T) {
	actual := Fizzbuzz(7)
	expected := "7"

	assert.Equal(t, expected, actual)
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByThree(t *testing.T) {
	actual := Fizzbuzz(3)
	expected := "Fizz"

	assert.Equal(t, expected, actual)
}

func TestFizzbuzzKnowsIfANumberIsNotDivisibleByThree(t *testing.T) {
	actual := Fizzbuzz(4)
	expected := "4"

	assert.Equal(t, expected, actual)
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByFive(t *testing.T) {
	actual := Fizzbuzz(5)
	expected := "Buzz"

	assert.Equal(t, expected, actual)
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByFifteen(t *testing.T) {
	actual := Fizzbuzz(15)
	expected := "Fizzbuzz"

	assert.Equal(t, expected, actual)
}
