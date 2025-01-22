package main

import (
	"fmt"
	"time"
)

func sample(item string) {
	for i := 0; i < 100; i++ {
		fmt.Println(item, ":", i)
	}
}

func main() {
	go sample("hello Borntodev")
	go sample("message2")
	time.Sleep(15 * time.Second)
}