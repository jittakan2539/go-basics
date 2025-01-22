package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)

	courseName = append(courseName, "Golang", "javaScript", "HTML", "CSS", "C")
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)
	
	courseWeb2 := courseName[:4]
	fmt.Println(courseWeb2)
}