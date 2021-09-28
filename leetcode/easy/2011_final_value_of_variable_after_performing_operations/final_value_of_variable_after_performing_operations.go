package main

import "fmt"

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}

func finalValueAfterOperations(operations []string) int {
	x := 0

	for _, op := range operations {
		switch op {
		case "X++", "++X":
			x++
		case "X--", "--X":
			x--
		}

	}
	return x
}
