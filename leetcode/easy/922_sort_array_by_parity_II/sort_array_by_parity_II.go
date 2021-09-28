package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	arr = sortArrayByParityII(arr)
	fmt.Println(arr)
}

func sortArrayByParityII(nums []int) []int {
	arr := make([]int, len(nums), len(nums))

	k := 0
	j := 1

	for _, v := range nums {
		if v%2 == 0 {
			arr[k] = v
			k += 2
		} else {
			arr[j] = v
			j += 2
		}
	}

	return arr
}
