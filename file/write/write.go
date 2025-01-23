package main

import "os"

func main() {
	data1 := []byte("hello\n my friend")
	err := os.WriteFile("../try.txt", data1, 0644)

	if err != nil {
		panic(err)
	}

	file, fileerr := os.Create("employeeName")
	if fileerr != nil {
		panic(fileerr)
	}

	defer file.Close()
	
	data2 := []byte("Wendy\n Rogers")
	os.WriteFile("employeeName", data2, 0644)

}