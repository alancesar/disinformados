package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	init := time.Now().UnixNano() / int64(time.Millisecond)
	text := readFile("./numbers.txt")
	numbers := parseText(text)
	results := process(numbers)

	for _, result := range results {
		fmt.Println(result)
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("Tempo total:", (end - init), "milissegundos")
}

func readFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(bytes)
}

func parseText(text string) []int {
	chars := strings.Split(text, "\n")
	numbers := []int{}

	for _, c := range chars {
		number, _ := strconv.Atoi(c)
		numbers = append(numbers, number)
	}

	return numbers
}

// (x² * √x) / (√x² + 1)
func calc(value int, channel chan float64) {
	channel <- (math.Pow(float64(value), 2) * math.Sqrt(float64(value))) / (math.Sqrt(math.Pow(float64(value), 2)) + 1)
}

func process(numbers []int) []float64 {
	results := []float64{}
	channel := make(chan float64)

	for _, number := range numbers {
		go calc(number, channel)
	}

	for {
		select {
		case result := <-channel:
			results = append(results, result)

			if len(results) == len(numbers) {
				return results
			}
		}
	}
}
