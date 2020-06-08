package main

import "fmt"

func main() {
	arr := [3]int{1: 10}
	arr[0] = 1
	fmt.Println(arr) // [1 10 0]

	arr = [...]int{2, 3, 5}
	fmt.Println(arr) // [2 3 5]
	// [3]int{2, 3, 5}
	fmt.Printf("%#v\n", arr)

	for i:= 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for index, value := range arr {
		fmt.Printf("index %d value %d\n", index, value)
	}
}
