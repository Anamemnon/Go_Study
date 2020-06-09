package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]

	var sum float64

	for _, value := range arguments {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}

		sum += number
	}

	fmt.Printf("Average: %0.2f\n", sum/(float64)(len(arguments)))
}

