package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	go printPares(channel)
	<-channel
	go printImpares()

	time.Sleep(time.Duration(500) * time.Millisecond)
}

func printPares(channel chan bool) {
	for i := 0; i < 100; i = i + 2 {
		fmt.Print(i, ", ")
	}

	channel <- true
}

func printImpares() {
	for i := 1; i < 100; i = i + 2 {
		fmt.Print(i, ", ")
	}
}
