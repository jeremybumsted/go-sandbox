package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Println("\nSimple Math Calculator")
		fmt.Println("---------------------")
		fmt.Println("Available operations: add, subtract, multiply, divide")
		fmt.Println("Enter 'exit' to quit")
		
		// Get first number
		fmt.Print("\nEnter first number: ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}
		a, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}
		
		// Get second number
		fmt.Print("Enter second number: ")
		if !scanner.Scan() {
			break
		}
		input = scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}
		b, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}
		
		// Get operation
		fmt.Print("Enter operation (add/subtract/multiply/divide): ")
		if !scanner.Scan() {
			break
		}
		operation := strings.ToLower(scanner.Text())
		if operation == "exit" {
			break
		}
		
		// Perform calculation
		result, err := Calculate(a, b, operation)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		
		// Display result
		fmt.Printf("\nResult: %.2f\n", result)
	}
	
	fmt.Println("Goodbye!")
}