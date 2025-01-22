package main

import "fmt"

func main() {
	// input := 4
	// switch input {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("Invalid Value")
	// } 
	
	var color string
	fmt.Scanf("%s", &color)
	switch color {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("Invalid color.")	
	}	
}