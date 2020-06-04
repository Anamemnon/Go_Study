package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readString()
	grade := parseNumber(input)
	fmt.Println(grade)
	fmt.Println("A grade of", grade, "is", getStatus(isPassedExam(grade)))
}

func readString() (input string) {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func parseNumber(parsingString string) (grade float64)  {
	parsingString = strings.TrimSpace(parsingString)
	grade, err := strconv.ParseFloat(parsingString, 64)
	if err != nil {
		log.Fatal(err)
	}
	return grade
}

func isPassedExam(grade float64) bool {
	return grade >= 60
}

func getStatus(passed bool) (status string) {
	if passed {
		return "passing"
	}
	return "failing"
}