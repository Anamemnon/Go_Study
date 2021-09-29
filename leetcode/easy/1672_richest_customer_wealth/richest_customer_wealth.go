package main

import "fmt"

func main() {
	fmt.Println(maximumWealth([][]int{{1,5}, {7,3}, {3,5}}))
}

func maximumWealth(accounts [][]int) int {
	max := 0

	for i := 0; i < len(accounts[0]); i++ {
		max += accounts[0][i]
	}

	for i := 1; i < len(accounts); i++ {
		temp := 0
		for j := 0; j < len(accounts[i]); j++ {
			temp += accounts[i][j]
		}

		if max < temp {
			max = temp
		}
	}

	return max
}