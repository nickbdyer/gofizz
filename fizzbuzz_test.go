package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzbuzzReturnNumberIfNotDivisible(t *testing.T) {
	assert.Equal(t, "7", Fizzbuzz(7))
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByThree(t *testing.T) {
	assert.Equal(t, "Fizz", Fizzbuzz(3))
}

func TestFizzbuzzKnowsIfANumberIsNotDivisibleByThree(t *testing.T) {
	assert.Equal(t, "4", Fizzbuzz(4))
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByFive(t *testing.T) {
	assert.Equal(t, "Buzz", Fizzbuzz(5))
}

func TestFizzbuzzKnowsIfANumberIsDivisibleByFifteen(t *testing.T) {
	assert.Equal(t, "Fizzbuzz", Fizzbuzz(15))
}
