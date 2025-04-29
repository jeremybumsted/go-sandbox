package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// generateRandomTestCases generates a specified number of random test cases
func generateRandomTestCases(count int) ([]float64, []float64) {
	rand.Seed(time.Now().UnixNano())
	a := make([]float64, count)
	b := make([]float64, count)
	
	for i := 0; i < count; i++ {
		// Generate random numbers between -100 and 100
		a[i] = (rand.Float64() * 200) - 100
		b[i] = (rand.Float64() * 200) - 100
	}
	
	return a, b
}

func TestAdd(t *testing.T) {
	// Fixed test cases
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 8.0},
		{"negative numbers", -2.0, -3.0, -5.0},
		{"mixed numbers", -5.0, 10.0, 5.0},
		{"zeros", 0.0, 0.0, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Add(test.a, test.b)
			assert.Equal(t, test.expected, result, "Addition failed for case: "+test.name)
		})
	}
	
	// Random test cases
	t.Run("random values", func(t *testing.T) {
		a, b := generateRandomTestCases(10)
		for i := 0; i < len(a); i++ {
			expected := a[i] + b[i]
			result := Add(a[i], b[i])
			assert.Equal(t, expected, result, "Addition failed for random case %d: %f + %f", i, a[i], b[i])
		}
	})
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 2.0},
		{"negative numbers", -2.0, -3.0, 1.0},
		{"mixed numbers", -5.0, 10.0, -15.0},
		{"zeros", 0.0, 0.0, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Subtract(test.a, test.b)
			assert.Equal(t, test.expected, result, "Subtraction failed for case: "+test.name)
		})
	}
	
	// Random test cases
	t.Run("random values", func(t *testing.T) {
		a, b := generateRandomTestCases(10)
		for i := 0; i < len(a); i++ {
			expected := a[i] - b[i]
			result := Subtract(a[i], b[i])
			assert.Equal(t, expected, result, "Subtraction failed for random case %d: %f - %f", i, a[i], b[i])
		}
	})
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 15.0},
		{"negative numbers", -2.0, -3.0, 6.0},
		{"mixed numbers", -5.0, 10.0, -50.0},
		{"zeros", 0.0, 5.0, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Multiply(test.a, test.b)
			assert.Equal(t, test.expected, result, "Multiplication failed for case: "+test.name)
		})
	}
	
	// Random test cases
	t.Run("random values", func(t *testing.T) {
		a, b := generateRandomTestCases(10)
		for i := 0; i < len(a); i++ {
			expected := a[i] * b[i]
			result := Multiply(a[i], b[i])
			assert.Equal(t, expected, result, "Multiplication failed for random case %d: %f * %f", i, a[i], b[i])
		}
	})
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
		wantErr  bool
	}{
		{"positive numbers", 6.0, 3.0, 2.0, false},
		{"negative numbers", -6.0, -3.0, 2.0, false},
		{"mixed numbers", -10.0, 2.0, -5.0, false},
		{"divide by zero", 5.0, 0.0, 0.0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Divide(test.a, test.b)
			
			if test.wantErr {
				assert.Error(t, err, "Expected error for division by zero")
			} else {
				assert.NoError(t, err, "Unexpected error during division")
				assert.Equal(t, test.expected, result, "Division failed for case: "+test.name)
			}
		})
	}
	
	// Random test cases
	t.Run("random values", func(t *testing.T) {
		a, b := generateRandomTestCases(10)
		for i := 0; i < len(a); i++ {
			// Skip any test cases where b is zero to avoid division by zero
			if b[i] == 0 {
				b[i] = 1.0 // Replace with non-zero value
			}
			
			expected := a[i] / b[i]
			result, err := Divide(a[i], b[i])
			
			assert.NoError(t, err)
			assert.Equal(t, expected, result, "Division failed for random case %d: %f / %f", i, a[i], b[i])
		}
	})
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		operation string
		expected  float64
		wantErr   bool
	}{
		{"addition", 5.0, 3.0, "add", 8.0, false},
		{"subtraction", 5.0, 3.0, "subtract", 2.0, false},
		{"multiplication", 5.0, 3.0, "multiply", 15.0, false},
		{"division", 6.0, 3.0, "divide", 2.0, false},
		{"division by zero", 5.0, 0.0, "divide", 0.0, true},
		{"unknown operation", 5.0, 3.0, "power", 0.0, true},
	}
	
	// Define all operations for random testing
	operations := []string{"add", "subtract", "multiply", "divide"}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Calculate(test.a, test.b, test.operation)
			
			if test.wantErr {
				assert.Error(t, err, "Expected error for case: "+test.name)
			} else {
				assert.NoError(t, err, "Unexpected error for case: "+test.name)
				assert.Equal(t, test.expected, result, "Calculation failed for case: "+test.name)
			}
		})
	}
	
	// Random test cases
	t.Run("random operations", func(t *testing.T) {
		a, b := generateRandomTestCases(20)
		
		for i := 0; i < len(a); i++ {
			// Choose a random operation
			opIndex := rand.Intn(len(operations))
			operation := operations[opIndex]
			
			// Skip division by zero cases
			if operation == "divide" && b[i] == 0 {
				b[i] = 1.0 // Replace with non-zero value
			}
			
			// Calculate expected result
			var expected float64
			switch operation {
			case "add":
				expected = a[i] + b[i]
			case "subtract":
				expected = a[i] - b[i]
			case "multiply":
				expected = a[i] * b[i]
			case "divide":
				expected = a[i] / b[i]
			}
			
			// Test the Calculate function
			result, err := Calculate(a[i], b[i], operation)
			
			assert.NoError(t, err)
			assert.Equal(t, expected, result, "Calculate failed for random case %d: %f %s %f", i, a[i], operation, b[i])
		}
	})
}

// TestFuzzyInputs tests the functions with very large, very small, and special values
func TestFuzzyInputs(t *testing.T) {
	// Special values to test
	specialValues := []float64{
		0.0,                  // Zero
		1.0,                  // One
		-1.0,                 // Negative one
		1e10,                 // Very large number
		1e-10,                // Very small number
		1e15,                 // Approaching float64 precision limits
		-1e15,                // Negative large number
	}
	
	for _, a := range specialValues {
		for _, b := range specialValues {
			// Skip division by zero
			if b == 0 {
				continue
			}
			
			t.Run("fuzzy_add", func(t *testing.T) {
				expected := a + b
				result := Add(a, b)
				assert.Equal(t, expected, result)
			})
			
			t.Run("fuzzy_subtract", func(t *testing.T) {
				expected := a - b
				result := Subtract(a, b)
				assert.Equal(t, expected, result)
			})
			
			t.Run("fuzzy_multiply", func(t *testing.T) {
				expected := a * b
				result := Multiply(a, b)
				assert.Equal(t, expected, result)
			})
			
			t.Run("fuzzy_divide", func(t *testing.T) {
				expected := a / b
				result, err := Divide(a, b)
				assert.NoError(t, err)
				assert.Equal(t, expected, result)
			})
		}
	}
}