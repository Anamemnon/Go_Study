package main

import (
	"fmt"
)

func main() {
	should_find := []rune("hello")
	max_index := len(should_find) - 1

	input := getRuns()
	index := 0

	for _, value := range input {
		if value == should_find[index] {
			if index == max_index {
				fmt.Println("YES")
				return
			}
			index++
		}
	}

	fmt.Println("NO")
}

func getRuns() []rune {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		panic(err)
	}
	return []rune(input)
}
