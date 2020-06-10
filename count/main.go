package main

import (
	"fmt"
	"github.com/Anamemnon/Go_Study/datafile"
	"log"
	"sort"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, value := range lines {
		counts[value]++
	}

	names := make([]string, 0)
	for name := range counts {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name, ":", counts[name])
	}
}
