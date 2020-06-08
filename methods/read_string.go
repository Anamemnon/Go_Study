package main

import (
	"Go_Study/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")

	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("A grade of", grade, "is", getStatus(isPassedExam(grade)))
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
