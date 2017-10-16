package main

import (
	"fmt"
	"time"
)

func main() {
	go printPares()
	go printImpares()

	time.Sleep(time.Duration(500) * time.Millisecond)
}

func printPares() {
	for i := 0; i < 100; i = i + 2 {
		fmt.Print(i, ", ")
	}
}

func printImpares() {
	for i := 1; i < 100; i = i + 2 {
		fmt.Print(i, ", ")
	}
}
