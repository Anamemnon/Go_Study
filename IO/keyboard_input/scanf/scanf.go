package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("name: ")
	name, err := GetName()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GetName():", name)

	fmt.Print("age: ")
	age, err := GetInt()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GetInt():", age)
}

func GetName() (string, error) {
	var name string
	_, err := fmt.Scanf("%s\n", &name)

	return name, err
}

func GetInt() (int, error) {
	var i int
	_, err := fmt.Scanf("%d\n", &i)

	return i, err
}
