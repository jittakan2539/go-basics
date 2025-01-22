package main

import (
	"fmt"
	"time"
)

func process1(c chan string, data string) {
	c <- data
}

func main() {
	channel := make(chan string)
	go process1(channel, "Data1")
	fmt.Println(<-channel)
	time.Sleep(5 * time.Second)
}