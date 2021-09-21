package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		txt := scanner.Text()

		if strings.Trim(txt, "\n\r ") == "" {
			break
		}

		fmt.Printf("Echo: %s\n", txt)
	}
}
