package main

import "fmt"

type subscriber struct {
	name        string
	rate        float64
	active      bool
	address
}

type address struct {
	street string
}

func main() {
	sub := subscriber{
		name:   "Aman",
		rate:   4.99,
		active: true,
		address: address{street: "Street"},
	}

	printSubscriber(&sub)
	fmt.Println(sub)
}

func printSubscriber(s *subscriber) {
	fmt.Println(s.name, s.rate, s.active, s.street)
	s.rate += 2
}
