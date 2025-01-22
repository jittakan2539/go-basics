package main

import "fmt"

var productName [4]string
// var price [4]float32


func array() {
	productName[0] = "Mcbook"
	productName[1] = "ipod"
	productName[2] = "ipad"
	productName[3] = "icloud"

	fmt.Println(productName)
}