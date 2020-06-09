package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64

	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}

		numbers = append(numbers, number)
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return numbers, nil
}
