package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{5, 0, 1, 2, 3, 4}))
}

func buildArray(nums []int) []int {
	arr := make([]int, 0, len(nums))

	for _, v := range nums {
		arr = append(arr, nums[v])
	}

	return arr
}
