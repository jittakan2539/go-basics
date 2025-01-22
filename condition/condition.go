package main

import "fmt"

func main() {
	var myScore int

	fmt.Println("Grade Calculator")
	fmt.Scanf("%d", &myScore)
	fmt.Println("score =", myScore)

	if myScore > 80 {
		fmt.Println("You got grade A")
	} else if myScore >= 70 {
		fmt.Println("You got grade B")
	} else if myScore >= 60 {
		fmt.Println("You got grade C")
	} else if myScore >= 50 {
		fmt.Println("You got grade D")
	} else {
		fmt.Println("You got grade F")
	}
	
}