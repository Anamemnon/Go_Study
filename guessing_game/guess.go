package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	targetNumber := getRandomNumber(100)

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	var guesses int

	for guesses = 0; guesses < 10; guesses++ {
		userNumber := getUserNumber()
		if isGuessed(targetNumber, userNumber) {
			break
		}
	}

	if guesses == 10 {
		println("Sorry. You didn't guess my number. It was", targetNumber)
	}

}

func getRandomNumber(upperBound int) int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	return rand.Intn(100) + 1 // from 1 to 100
}

func getUserNumber() int {
	fmt.Print("Your number: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	return number
}

func isGuessed(target, guess int) bool {
	if guess < target {
		fmt.Println("Your guess was LOW.")
	} else if guess > target {
		fmt.Println("Your guess was HIGH.")
	} else {
		fmt.Println("You guessed !")
		return true
	}
	return false
}
