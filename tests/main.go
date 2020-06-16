package main

func main() {

}

func sum(numbers ...int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}