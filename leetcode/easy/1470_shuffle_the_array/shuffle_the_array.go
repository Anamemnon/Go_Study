package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{2,5,1,3,4,7}, 3))
}

func shuffle(nums []int, n int) []int {
	arr := make([]int, 0, n * 2)

	for i := 0; i < n; i++ {
		arr = append(arr, nums[i])
		arr = append(arr, nums[n + i])
	}

	return arr
}