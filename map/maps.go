package main

import "fmt"



func main() {
	product := make(map[string]float64)
	fmt.Println("product", product)

	//add
	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 25000
	fmt.Println("product", product)

	//delete
	delete(product, "iPad")
	fmt.Println("product", product)

	//update
	product["iPhone"] = 20000
	fmt.Println("product", product)

	//access data
	value1 := product["iPhone"]
	fmt.Println("value1", value1)

	// -------------------------//
	courseName := map[string]string{"101":"Java","102":"Python"}
	fmt.Println("courseName", courseName)
}