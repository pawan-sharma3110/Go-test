package main

import (
	"fmt"
	"os"
)

func add(a, b float64) float64 {
	return a + b
}
func subtract(a, b float64) float64 {
	return a - b
}
func multiply(a, b float64) float64 {
	return a * b
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
func calculator() {
	var num1, num2 float64
	var operator string
	fmt.Println("Select two numbers and an operator from (+, -, *, /)")
	fmt.Print("Enter your first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter your operator: ")
	fmt.Scanln(&operator)
	oprator := rune(operator[0]) // Convert the first character of the operator string to rune
	fmt.Print("Enter your second number: ")
	fmt.Scanln(&num2)
	var result float64
	var err error
	switch oprator {
	case '+':
		result = add(num1, num2)
	case '-':
		result = subtract(num1, num2)
	case '*':
		result = multiply(num1, num2)
	case '/':
		result, err = divide(num1, num2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Select the right operator")
		os.Exit(1)
	}
	fmt.Println("Result:", result)
}
