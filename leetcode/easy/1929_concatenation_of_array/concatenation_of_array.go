package main

import "fmt"

func main() {
	arr := []int{1, 2}
	fmt.Println(getConcatenation((arr)))
}

func getConcatenation(nums []int) []int {
	arr := make([]int, 0, len(nums)*2)
	arr = append(arr, nums...)
	arr = append(arr, nums...)
	return arr
}
