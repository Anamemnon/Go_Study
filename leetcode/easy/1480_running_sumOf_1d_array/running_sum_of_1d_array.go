package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1,2,3,4}))
}

func runningSum(nums []int) []int {
	sum := nums[0]

	for i:=1 ; i < len(nums) ; i++ {
		sum += nums[i]
		nums[i] = sum
	}
	return nums
}