package main

import "fmt"

var productName [4]string
// var price [4]float32


func main() {
	productName[0] = "Mcbook"
	productName[1] = "ipod"
	productName[2] = "ipad"
	productName[3] = "icloud"

	price := [4]float32{4000,3000,2000,2500}

	fmt.Println(productName)
	fmt.Println(price)
}