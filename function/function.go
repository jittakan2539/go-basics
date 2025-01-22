package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func plus(value1, value2 int) int {
	result := value1 + value2
	return result
}

func main() {
	hello()
	plus(1,2)
	returnResult := plus(1,2)
	fmt.Println("returnResult", returnResult)
}