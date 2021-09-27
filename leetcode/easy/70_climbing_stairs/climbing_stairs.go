package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	items := [3]int{0, 1, 2}

	for i := 3; i <= n; i++ {
		items[0], items[1] = items[1], items[2]
		items[2] = items[0] + items[1]
	}
	if n < 3 {
		return items[n]
	}

	return items[2]
}
