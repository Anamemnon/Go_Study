package main

import (
	"fmt"
	"reflect"
)

func main() {
	var amount int = 5

	fmt.Println(amount)

	fmt.Println(&amount)
	fmt.Println(reflect.TypeOf(&amount))

	var pAmount *int = &amount
	*pAmount = 3

	fmt.Println(amount)

	doubleValue(&amount)
	fmt.Println(amount)
}

func doubleValue(n *int) {
	*n *= 2
}
