package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}