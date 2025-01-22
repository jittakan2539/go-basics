package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must be a number only.", prompt)
		panic(message) 
	}

	return value
}
func getInteger(prompt string) int {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must be an integer only.", prompt)
		panic(message) 
	}

	return int(value)
}

func getOperator() string {
	fmt.Println("Operators are ( + - * /):")
	operator, _ := reader.ReadString('\n')
	return strings.TrimSpace((operator))
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}
func subtract(value1, value2 float64) float64 {
	return value1 - value2
}
func multiply(value1, value2 float64) float64 {
	return value1 * value2
}
func divide(value1, value2 float64) float64 {
	if value2 == 0 {
		panic(fmt.Sprintf("Division by zero is not allowed: value2 = %v: ", value2))
	}
	return value1 / value2
}

func main() {
	count := getInteger(" How many numbers would you like to compute? ")

	//Create slices for number and operators
	values := make([]float64, count)
	operators := make([]string, count-1)

	//initial value
	values[0] = getInput("Number 1: ")

	for i := 1; i < count; i++ {
		fmt.Println("Hello")
		operators[i-1] = getOperator()
		values[i] = getInput(fmt.Sprintf("Number %d: ", i + 1))
	}

	//start with initial value
	result := values[0]

	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result = add(result, values[i+1])
		case "-":
			result = subtract(result, values[i+1])
		case "*":
			result = multiply(result, values[i+1])
		case "/":
			result = divide(result, values[i+1])
		default:
			panic("Invalid Operator")		
		}
	}
	
	fmt.Printf("The result is %v", result)
}