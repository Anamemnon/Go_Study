package main

import (
	"fmt"
	"reflect"
)

func main()  {
	amount := 5
	fmt.Println(amount)
	fmt.Println(&amount)
	fmt.Println(reflect.TypeOf(&amount))

	p_amount := &amount
	*p_amount = 3
	fmt.Println(amount)

	double_value(&amount)
	fmt.Println(amount)
}

func double_value(n *int) {
	*n *= 2
}