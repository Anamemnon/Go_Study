package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	numbers := make([]float64, 0)

	for _, value := range arguments {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	fmt.Printf("Average: %0.2f\n", average(numbers...))
}

func average(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	var sum float64
	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}
