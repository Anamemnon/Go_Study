package main

import (
	"fmt"
	"log"
)

func main() {
	number, err := getFloat()

	if err != nil {
		log.Fatal(err)
	}

	switch {
	case number <= 0:
		fmt.Printf("число %2.2f не подходит\n", number)
		return
	case number > 10000:
		fmt.Printf("%e\n", number)
		return
	}

	squareNumber := number * number
	fmt.Printf("%2.4f", squareNumber)
}

func getFloat() (float64, error) {
	var number float64
	_, err := fmt.Scanf("%f", &number)

	if err != nil {
		return 0, err
	}

	return number, nil
}
