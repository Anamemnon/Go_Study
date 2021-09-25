package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := getInt()
	scanner := bufio.NewScanner(os.Stdin)
	passenger := passengers{}
	max := 0
	temp := 0

	for i := 0; i < n; i++ {
		str := nextString(scanner)

		if str == "" {
			break
		}

		parse(str, &passenger)
		temp = temp - passenger.Out + passenger.In

		if temp > max {
			max = temp
		}
	}

	fmt.Println(max)

}

type passengers struct {
	In  int
	Out int
}

func getInt() int {
	var i int
	_, err := fmt.Scanf("%d\n", &i)

	if err != nil {
		panic(err)
	}

	return i
}

func nextString(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func parse(input string, p *passengers) {

	arr := strings.Split(input, " ")
	out, err := strconv.Atoi(arr[0])

	if err != nil {
		panic(err)
	}

	in, err := strconv.Atoi(arr[1])

	if err != nil {
		panic(err)
	}

	p.In = in
	p.Out = out
}
