package main

import (
	"errors"
	"fmt"
)

// Add adds two numbers and returns the result
func Add(a, b float64) float64 {
	return a + b
}

// Subtract subtracts b from a and returns the result
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply multiplies two numbers and returns the result
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide divides a by b and returns the result
// Returns an error if division by zero is attempted
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Calculate performs a mathematical operation on two numbers
// Supported operations: add, subtract, multiply, divide
func Calculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "add":
		return Add(a, b), nil
	case "subtract":
		return Subtract(a, b), nil
	case "multiply":
		return Multiply(a, b), nil
	case "divide":
		return Divide(a, b)
	default:
		return 0, fmt.Errorf("unsupported operation: %s", operation)
	}
}