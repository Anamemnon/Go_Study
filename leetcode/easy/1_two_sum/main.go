package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	var numsIndx = make(map[int]int)

	for i, n := range nums {
		if firstIndex, ok := numsIndx[target-n]; ok {
			return []int{firstIndex, i}
		}

		numsIndx[n] = i
	}
	return []int{}
}
