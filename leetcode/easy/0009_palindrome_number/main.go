package main

import (
	"strconv"
)

func main() {
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	strX := strconv.Itoa(x)
	length := len(strX)
	for i, v := range strX {
		if string(v) != string(strX[length-1-i]) {
			return false
		}
	}
	return true
}
