package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.OpenFile("test.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner()
}